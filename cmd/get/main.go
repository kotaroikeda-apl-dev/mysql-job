package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "user:password@tcp(localhost:3306)/job_db"

func getJobLogs() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, job_name, status, retry_count, created_at FROM job_logs")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id         int
			jobName    string
			status     string
			retryCount int
			createdAt  string
		)
		if err := rows.Scan(&id, &jobName, &status, &retryCount, &createdAt); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Job: %s, Status: %s, Retries: %d, Created: %s\n",
			id, jobName, status, retryCount, createdAt)
	}
}

func main() {
	getJobLogs()
}
