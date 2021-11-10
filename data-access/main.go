package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type Record struct {
	ID               int64          `json:"id"`
	AppId            int64          `json:"app_id"`
	DeviceSn         string         `json:"device_sn"`
	UserId           int64          `json:"user_id"`
	EntityName       string         `json:"entity_name"`
	AttributesDigest string         `json:"attributes_digest"`
	CreatedAt        string         `json:"created_at"`
	UpdatedAt        string         `json:"updated_at"`
	DeletedAt        sql.NullString `json:"deleted_at"`
}

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "192.168.86.129:3306",
		DBName:               "go",
		AllowNativePasswords: true, // 默认是false，会报：could not use requested auth plugin 'mysql_native_password': this user requires mysql native password authentication.
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	//// 查询多行
	//records, err := getRecordsByDeviceSn("e7f15a647b10b04d933cf74953a6d6fbbfb64250")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("records found：%v\n", records)

	//recordId, err := addRecord(Record{
	//	AppId:            10010,
	//	DeviceSn:         "e7f15a647b10b04d933cf74953a6d6fbbfb64250",
	//	UserId:           0,
	//	EntityName:       "qiming",
	//	AttributesDigest: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
	//	CreatedAt:        time.Now().Format("2006-1-2 15:04:05"),
	//	UpdatedAt:        time.Now().Format("2006-1-2 15:04:05"),
	//	DeletedAt:        sql.NullString{"", false},
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("insert new record：%v\n", recordId)

	//// 查询单行
	//record, err := getInfoById(1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Record found: %v\n", record)
}

/**
根据设备号查询
*/
func getRecordsByDeviceSn(deviceSn string) ([]Record, error) {
	var records []Record
	rows, err := db.Query("select * from records where device_sn=?", deviceSn)
	if err != nil {
		return nil, fmt.Errorf("getRecordsByDeviceSn %q：%v", deviceSn, err)
	}
	// 关闭连接
	defer rows.Close()

	for rows.Next() {
		var rec Record
		if err := rows.Scan(&rec.ID, &rec.AppId, &rec.DeviceSn, &rec.UserId, &rec.EntityName, &rec.AttributesDigest, &rec.CreatedAt, &rec.UpdatedAt, &rec.DeletedAt); err != nil {
			return nil, fmt.Errorf("getRecordsByDeviceSn %q：%v", deviceSn, err)
		}
		records = append(records, rec)
	}
	return records, nil
}

/**
新增数据
*/
func addRecord(rec Record) (int64, error) {
	//result, err := db.Exec(
	//	"insert into records (`app_id`, `device_sn`, `user_id`, `entity_name`, `attributes_digest`, `created_at`, `updated_at`) VALUES (?,?,?,?,?,?,?)",
	//	rec.AppId, rec.DeviceSn, rec.UserId, rec.EntityName, rec.AttributesDigest, rec.CreatedAt, rec.UpdatedAt,
	//)

	result, err := db.Exec(
		"INSERT INTO records (app_id, device_sn, user_id,entity_name,attributes_digest,created_at,updated_at,deleted_at) VALUES (?, ?, ?,?, ?, ?,?,?)",
		rec.AppId, rec.DeviceSn, rec.UserId, rec.EntityName, rec.AttributesDigest, rec.CreatedAt, rec.UpdatedAt, rec.DeletedAt)
	if err != nil {
		return 0, fmt.Errorf("addRecord：%v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addRecord：%d", err)
	}
	return id, nil
}

/**
根据id查询
*/
func getInfoById(id int64) (Record, error) {
	var rec Record

	row := db.QueryRow("SELECT * FROM records WHERE id = ?", id)
	if err := row.Scan(&rec.ID, &rec.AppId, &rec.DeviceSn, &rec.UserId, &rec.EntityName, &rec.AttributesDigest, &rec.CreatedAt, &rec.UpdatedAt, &rec.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return rec, fmt.Errorf("albumsById %d: no such album", id)
		}
		return rec, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return rec, nil
}

/**
插入档案事务
*/
func InsertTransaction(ctx context.Context, rec Record) (int64, error) {
	//
	fail := func(err error) (int64 error) {
		return fmt.Errorf("addRecord：%d", err)
	}
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fail(err)
	}
	defer tx.Rollback()

	// 插入档案
	result, err := db.Exec(
		"INSERT INTO records (app_id, device_sn, user_id,entity_name,attributes_digest,created_at,updated_at,deleted_at) VALUES (?, ?, ?,?, ?, ?,?,?)",
		rec.AppId, rec.DeviceSn, rec.UserId, rec.EntityName, rec.AttributesDigest, rec.CreatedAt, rec.UpdatedAt, rec.DeletedAt)
	if err != nil {
		return 0, fail(err)
	}
	recordId, err := result.LastInsertId()
	if err != nil {
		return 0, fail(err)
	}

	// 插入档案拓展表
	recordExtension, err := db.Exec("INSERT INTO record_extension (record_id, xiyongshen_version) VALUES (?, ?)", recordId, 1)

	if err != nil {
		return 0, fail(err)
	}
	_, err = recordExtension.LastInsertId()
	if err != nil {
		return 0, fail(err)
	}
	if err = tx.Commit(); err != nil {
		return 0, fail(err)
	}
	return recordId, nil
}
