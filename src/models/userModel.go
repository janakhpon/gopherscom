package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID        string    `json:"_id" json:"id"`
	NAME      string    `json:"name"`
	EMAIL     string    `json:"email"`
	PASSWORD  string    `json:"password"`
	CREATEDAT time.Time `json:"created_at" json:"created_at"`
	UPDATEDAT time.Time `json:"updated_at" json:"updated_at"`
}

type JWT struct {
	Token string `json:"token"`
}

type Claims struct {
	ID        string    `json:"_id" json:"id"`
	NAME      string    `json:"name"`
	EMAIL     string    `json:"email"`
	CREATEDAT time.Time `json:"created_at" json:"created_at"`
	UPDATEDAT time.Time `json:"updated_at" json:"updated_at"`
	jwt.StandardClaims
}
