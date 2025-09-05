package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(ds *sql.DB) {
	http.HandleFunc("/create", CreateHandler(ds))
}