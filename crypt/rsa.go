package crypt

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func NewPrivateKey() *rsa.PrivateKey {
	if privateKey == nil {
		key, err := createPrivateKey()
		if err != nil {
			panic(err)
		}
		privateKey = key
	}
	return privateKey
}
func NewPublicKey() *rsa.PublicKey {
	if publicKey == nil {
		key, err := createPublicKey()
		if err != nil {
			panic(err)
		}
		publicKey = key
	}
	return publicKey
}
func createPublicKey() (key *rsa.PublicKey, err error) {
	// 公開鍵を読み込む
	pubKeyData, err := ioutil.ReadFile("./rsa/public-key.pem")
	if err != nil {
		return
	}
	key, err = jwt.ParseRSAPublicKeyFromPEM(pubKeyData)
	return
}
func createPrivateKey() (key *rsa.PrivateKey, err error) {
	// 秘密鍵を読み込み
	privateKeyData, err := ioutil.ReadFile("./rsa/private-key.pem")
	if err != nil {
		return
	}

	key, err = jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	return
}
