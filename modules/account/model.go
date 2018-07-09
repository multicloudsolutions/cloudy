package account

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	//mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Account type provides the basic structure for inputting various cloud provider details
type Account struct {
	gorm.Model
	Provider  string `json:"provider" gorm:"unique;not null"`
	Name      string `json:"name" gorm:"unique;not null"`
	AccessKey string `json:"access_key" gorm:"unique;not null"`
	SecretKey string `json:"secret_key" gorm:"unique;not null"`
}

// GetAccount Get details of and account by name
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

// CreateAccount Enter the details of a new account
func (acc Account) CreateAccount() string {
	db, err := gorm.Open("mysql", "test:test@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Create(&acc)
	return "Created account"
}

// RemoveAccount by name
func RemoveAccount(name string) string {
	db, err := gorm.Open("mysql", "test:test@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Where("name = ?", name).Delete(&Account{})
	return "Deleted account"

}
