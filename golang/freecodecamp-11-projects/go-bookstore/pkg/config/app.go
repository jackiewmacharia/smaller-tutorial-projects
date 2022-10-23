package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// DB_CONNECTION_URL example: export DB_CONNECTION_URL="USERNAME:PASSWORD@tcp(HOST:PORT)/DBNAME?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", os.Getenv("DB_CONNECTION_URL"))
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
