package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *sql.DB

type Record struct {
	ID               int64  `json:"id"`
	AppId            int64  `json:"app_id"`
	DeviceSn         string `json:"device_sn"`
	UserId           int64  `json:"user_id"`
	EntityName       string `json:"entity_name"`
	AttributesDigest string `json:"attributes_digest"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

func main() {
	// Capture connection properties.
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

	records, err := getRecordsByDeviceSn("aaa")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("records found：%v\n", records)

	recordId, err := addRecord(Record{
		AppId:            10010,
		DeviceSn:         "aaa",
		UserId:           123,
		EntityName:       "qiming",
		AttributesDigest: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		CreatedAt:        time.Now().Format("2006-1-2 15:04:05"),
		UpdatedAt:        time.Now().Format("2006-1-2 15:04:05"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("insert new record：%v\n", recordId)

	record, err := getInfoById(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Record found: %v\n", record)
}

/**
根据设备号查询
*/
func getRecordsByDeviceSn(deviceSn string) ([]Record, error) {
	var records []Record
	//db, _ := connection()
	rows, err := db.Query("select * from records where device_sn=?", deviceSn)
	if err != nil {
		return nil, fmt.Errorf("getRecordsByDeviceSn %q：%v", deviceSn, err)
	}
	// 关闭连接
	defer rows.Close()

	for rows.Next() {
		var rec Record
		if err := rows.Scan(&rec.ID, &rec.AppId, &rec.DeviceSn, &rec.UserId, &rec.EntityName, &rec.AttributesDigest, &rec.CreatedAt, &rec.UpdatedAt); err != nil {
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
	result, err := db.Exec("INSERT INTO records (app_id, device_sn, user_id,entity_name,attributes_digest,created_at,updated_at) VALUES (?, ?, ?,?, ?, ?,?)", rec.AppId, rec.DeviceSn, rec.UserId, rec.EntityName, rec.AttributesDigest, rec.CreatedAt, rec.UpdatedAt)
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
	if err := row.Scan(&rec.ID, &rec.AppId, &rec.DeviceSn, &rec.UserId, &rec.EntityName, &rec.AttributesDigest, &rec.CreatedAt, &rec.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return rec, fmt.Errorf("albumsById %d: no such album", id)
		}
		return rec, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return rec, nil
}
