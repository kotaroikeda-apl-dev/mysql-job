package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "user:password@tcp(localhost:3306)/job_db"

func updateJobStatus(id int, status string, retryCount int) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	query := `UPDATE job_logs SET status = ?, retry_count = ? WHERE id = ?`
	_, err = db.Exec(query, status, retryCount, id)
	return err
}

func main() {
	err := updateJobStatus(1, "SUCCESS", 0)
	if err != nil {
		log.Fatal("ジョブの更新に失敗:", err)
	} else {
		fmt.Println("ジョブの更新成功")
	}
}
