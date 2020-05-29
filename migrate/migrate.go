package main

import (
	"fmt"
	"os"
	"golang.org/x/crypto/bcrypt"

	"github.com/yamakenji24/binder-api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db_user string
	db_pass string
	db_host string
	db_name string
)

func main() {
	db_user = os.Getenv("DB_USER")
	db_pass = os.Getenv("DB_PASSWORD")
	db_name = os.Getenv("DB_NAME")
	db_host = os.Getenv("DB_HOST")
	param := db_user + ":" + db_pass + "@(" + db_host + ")/" + db_name + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", param)
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Successful connection")
	}

	db.AutoMigrate(&model.User{})

	sampleUser := model.User{
		Username:     "yamakenji24`",
		Password: toHashPassword("yamakenji24"),
		Email:    "yamakenji24@example.com",
	}
	db.Create(&sampleUser)
}


func toHashPassword(pass string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(converted)
}
