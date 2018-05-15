package main

import (
	"github.com/cloudy/internal/app/dummy_app"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/", dummy_app.MyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
