package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/multicloudsolutions/cloudy/modules/config"
)

// DBHandler encapsulates gorm.DB
type DBHandler struct {
	Db *gorm.DB
}

//GetConn gets a db commection to gorm.DB
func (hdlr *DBHandler) GetConn() {
	var err error
	connectionString := GetDBConnectionString()
	hdlr.Db, err = gorm.Open("mysql", connectionString)
	GetDBConnectionString()
	if err != nil {
		panic(err)
	}
}

// GetDBConnectionString to get connection string for db
func GetDBConnectionString() string {
	var dbConf config.DBConfig
	var connectionString string
	(&dbConf).GetConfig()
	connectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		dbConf.DBUsername,
		dbConf.DBPassword,
		dbConf.DBName,
	)

	return connectionString
}
