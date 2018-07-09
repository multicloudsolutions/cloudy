package utils

import (
	"github.com/jinzhu/gorm"
)

type DBHandler struct {
	Db *gorm.DB
}

func (hdlr *DBHandler) GetConn() {
	var err error
	hdlr.Db, err = gorm.Open("mysql", "test:test@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
