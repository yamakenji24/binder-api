package repository

import (
	"fmt"

	"github.com/yamakenji24/binder-api/model"
)

func GetUserByName(username string) (user model.User, err error) {
	db := NewSQLHandler()
	if res := db.Where("username = ?", username).Find(&user); res.Error != nil {
		err = fmt.Errorf("error in GetUserByName() : %w", res.Error)
		return
	}
	return
}
