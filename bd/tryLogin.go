package bd

import (
	"Documents/Go/twitter-go/models"

	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, passwd string) (models.User, bool) {
	usu, finded, _ := CheckUserExists(email)
	if finded == false {
		return usu, false
	}

	passwdBytes := []byte(passwd)
	passwdBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwdBD, passwdBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
