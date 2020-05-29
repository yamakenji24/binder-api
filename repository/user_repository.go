package repository

import (
	"github.com/yamakenji24/binder-api/model"
)

func GetUserByName(username string) (*model.User, error) {
	return &model.User{
		Username: username,
		Password: "password",
		Email:    "yamakenji24@example.com",
	}, nil
}
