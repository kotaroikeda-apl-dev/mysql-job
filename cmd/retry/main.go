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

func retryFailedJobs() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id FROM job_logs WHERE status = 'FAILED' AND retry_count < 3")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		if err := updateJobStatus(id, "RETRYING", 1); err != nil {
			log.Printf("ジョブ %d の更新に失敗: %v\n", id, err)
		} else {
			fmt.Printf("ジョブ %d をリトライ\n", id)
		}
	}

	// `FAILED` かつ `retry_count >= 3` のジョブを取得
	rows, err = db.Query("SELECT id, job_name, retry_count FROM job_logs WHERE status = 'FAILED' AND retry_count >= 3")
	if err != nil {
		log.Fatal("クエリ実行に失敗:", err)
	}
	defer rows.Close()
	// 結果を出力
	found := false
	for rows.Next() {
		var id, retryCount int
		var jobName string

		if err := rows.Scan(&id, &jobName, &retryCount); err != nil {
			log.Fatal("データ取得に失敗:", err)
		}

		fmt.Printf("ジョブID: %d, ジョブ名: %s, リトライ回数: %d (リトライ上限到達)\n", id, jobName, retryCount)
		found = true
	}

	// `FAILED` かつ `retry_count >= 3` のジョブがない場合
	if !found {
		fmt.Println("リトライ上限に達した `FAILED` のジョブはありません。")
	}
}

func main() {
	retryFailedJobs()
}
