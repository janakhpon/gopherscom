package utils

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/janakhpon/gopherscom/src/models"
)

func GenerateToken(user models.User) (string, string, error) {
	var err error
	secret := os.Getenv("SECRET")

	expirationTime := time.Now().Add(time.Hour * 2)
	reexpirationTime := time.Now().Add(time.Hour * 240)

	claims := &models.Claims{
		ID:        user.ID,
		EMAIL:     user.EMAIL,
		NAME:      user.NAME,
		UPDATEDAT: user.UPDATEDAT,
		CREATEDAT: user.CREATEDAT,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	reclaims := &models.Claims{
		EMAIL: user.EMAIL,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: reexpirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, reclaims)
	tokenString, err := token.SignedString([]byte(secret))
	retokenString, err := refreshtoken.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, retokenString, nil
}
