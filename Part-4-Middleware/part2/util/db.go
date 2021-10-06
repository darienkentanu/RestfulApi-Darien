package util

import (
	"fmt"
	"part2/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// func init() {
// 	InitDB()
// 	InitialMigration()
// }

// var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {
	config := Config{"root", "password", "3306", "localhost", "crud_go"}
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username, config.DB_Password, config.DB_Host, config.DB_Port, config.DB_Name,
	)
	var err error
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&model.Book{})
}
