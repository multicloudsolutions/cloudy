package main

import (
	"log"
    "net/http"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
    "github.com/multicloudsolutions/cloudy/modules"

)

func main() {
	db, err := sql.Open("mysql", "test:test@/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	a := ac.AccountHandler{Db: db}
    http.Handle("/account", &a)
    http.ListenAndServe(":8080", nil)
}
