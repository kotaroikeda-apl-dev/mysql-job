package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "user:password@tcp(localhost:3306)/job_db"

func insertJobLog(jobName, status string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO job_logs (job_name, status) VALUES (?, ?)`
	_, err = db.Exec(query, jobName, status)
	return err
}

func main() {
	err := insertJobLog("data_processing", "PENDING")
	if err != nil {
		log.Fatal("ジョブログの登録に失敗:", err)
	} else {
		fmt.Println("ジョブログ登録成功")
	}
}
