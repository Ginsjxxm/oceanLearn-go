package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"wenzhang/model"
)

var jwtKey = []byte("a_secret_key")

type Claims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "oceanlearn.tech",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, Claims, err
}
