package main

// This is a test comment

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/multicloudsolutions/cloudy/modules/account"
	"net/http"
)

func main() {
	db, err := gorm.Open("mysql", "test:test@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print(err)
	}
	db.AutoMigrate(&account.Account{})
	db.Close()

	accountRouter := account.Router()
	http.Handle("/account", accountRouter)
	http.ListenAndServe(":8080", accountRouter)
}
