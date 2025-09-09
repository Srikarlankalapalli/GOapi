package main

import (
	"GOapi/api"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:Srikar@1930@tcp(127.0.0.1:3306)/library_db?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	api.RegisterRoutes(db)

	log.Println("Server Staring on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
