package config

import (
	"encoding/json"
	"os"
)

//DBConfig for parsing database config
type DBConfig struct {
	DBName     string `json:"db_name"`
	DBPassword string `json:"db_password"`
	DBUsername string `json:"db_username"`
	DBType     string `json:"db_type"`
}

// GetConfig to get dabase config
func (DBConf *DBConfig) GetConfig() {
	file, err := os.Open("modules/config/db.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&DBConf)
	if err != nil {
		panic(err)
	}
}
