package repository

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

var (
	db_user string
	db_pass string
	db_host string
	db_name string
)

func NewSQLHandler() *gorm.DB {
	if db != nil {
		return db
	}

	db_user = os.Getenv("DB_USER")
	db_pass = os.Getenv("DB_PASSWORD")
	db_name = os.Getenv("DB_NAME")
	db_host = os.Getenv("DB_HOST")
	param := db_user + ":" + db_pass + "@(" + db_host + ")/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", param)
	if err != nil {
		panic(err)
	}
	return db
}
