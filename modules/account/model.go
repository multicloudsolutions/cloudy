package account

import (
      "fmt"
      "encoding/json"
      "github.com/jinzhu/gorm"
      _ "github.com/jinzhu/gorm/dialects/mysql"
)


type Account struct {
	gorm.Model
        Provider string `json:"provider" gorm:"unique;not null"`
        Name string `json:"name" gorm:"unique;not null"`
        AccessKey string `json:"access_key" gorm:"unique;not null"`
        SecretKey string `json:"secret_key" gorm:"unique;not null"`
}


func GetAccount(name string) string {
        db, err := gorm.Open("mysql", "test:test@/test?charset=utf8&parseTime=True&loc=Local")
        if err != nil {
            fmt.Print(err)
        }
	defer db.Close()
        var ac []Account
        db.First(&ac, "name=?", name)
        res, err := json.Marshal(&ac)
        if err != nil {
               panic(err)
        }

        return string(res)
}


func (acc Account) CreateAccount() string {
        db, err := gorm.Open("mysql", "test:test@/test?charset=utf8&parseTime=True&loc=Local")
        if err != nil {
            fmt.Print(err)
        }
	defer db.Close()
        db.Create(&acc)
        return "Created account"
}
