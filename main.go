package main

import (
	"GOapi/api"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dns := "user=srikar password @tcp(localhost(127.0.0.1)dbname?parseTime=true"

	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Error opening database: ", err)

	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection established successfully!")

	api.RegisterRoutes(db)

	log.Println("Server started on :3030")
	log.Fatal(http.ListenAndServe(":3030", nil))
}