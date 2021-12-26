package jwt

import (
	"Documents/Go/twitter-go/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.User) (string, error) {

	myKey := []byte("MiClave")
	payload := jwt.MapClaims{
		"email":    user.Email,
		"name":     user.Name,
		"lastname": user.LastName,
		"birthday": user.Birthday,
		"bio":      user.Bio,
		"_id":      user.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
