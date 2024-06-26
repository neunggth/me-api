package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		log.Fatal("Can't prepare delete statment", err)
	}
	if _, err = stmt.Exec(1); err != nil {
		log.Fatal("can't excute delete statment", err)
	}

	fmt.Println("Delete success! ")

}
