package routers

import (
	"Documents/Go/twitter-go/bd"
	"Documents/Go/twitter-go/models"
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUser string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("MiClave")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, claims, func(toke *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		_, finded, _ := bd.CheckUserExists(claims.Email)
		if finded == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, finded, IDUser, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}

	return claims, false, string(""), err
}
