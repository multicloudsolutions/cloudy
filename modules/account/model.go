package account

import (
	//	"encoding/json"
	"github.com/jinzhu/gorm"
	//mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/multicloudsolutions/cloudy/modules/utils"
)

// Account type provides the basic structure for inputting various cloud provider details
type Account struct {
	gorm.Model
	Provider  string `json:"provider" gorm:"not null"`
	Name      string `json:"name" gorm:"unique;not null"`
	AccessKey string `json:"access_key" gorm:"not null"`
	SecretKey string `json:"secret_key" gorm:"not null"`
}

// GetAccount Get details of and account by name
func (ac *Account) GetAccount(name string) {
	var dbHandler utils.DBHandler
	(&dbHandler).GetConn()
	defer dbHandler.Db.Close()
	dbHandler.Db.First(&ac, "name=?", name)
}

// CreateAccount Enter the details of a new account
func (ac *Account) CreateAccount() string {
	var dbHandler utils.DBHandler
	(&dbHandler).GetConn()
	defer dbHandler.Db.Close()
	dbHandler.Db.Create(ac)
	return "Created account"
}

// RemoveAccount by name
func (ac *Account) RemoveAccount(name string) string {
	var dbHandler utils.DBHandler
	(&dbHandler).GetConn()
	dbHandler.Db.Where("name = ?", name).Delete(&Account{})
	return "Deleted account"
}

//// UpdateAccount updates account details
//func (acc Account) UpdateAccount() string {
//	db, err := gorm.Open("mysql", "test:test@/test?charset=utf8&parseTime=True&loc=Local")
//	if err != nil {
//		fmt.Print(err)
//	}
//	defer db.Close()
//	result := GetAccount(acc.name)
//	db.Model(&Account{}).Updates(acc)
//	return "Updated account"
//}
