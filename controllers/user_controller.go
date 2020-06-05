package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yamakenji24/binder-api/crypt"
	"github.com/yamakenji24/binder-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type InputUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) error {
	input := &InputUser{}

	if err := c.BindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	user, err := repository.GetUserByName(input.Username)
	if err != nil {
		return err
	}
	if !compareHashedPassword(user.Password, input.Password) {
		return fmt.Errorf("Invalid username and password combination!")
	}
	// TODO: create jwt token
	key := crypt.NewPrivateKey()
	token := jwt.New(jwt.SigningMethodRS256)
	// claimの設定
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["name"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 60).Unix()

	t, err := token.SignedString(key)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}{
		Status: "success",
		Token:  t,
	})
	return nil
}

func compareHashedPassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return true
	}
	return false
}
