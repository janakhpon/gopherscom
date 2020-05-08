package models

import (
	"time"
)

type Profile struct {
	ID         string    `json:"id"`
	USER       string    `json:"user"`
	CAREER     string    `json:"career"`
	LANGUAGES  []string  `json:"languages"`
	FRAMEWORKS []string  `json:"frameworks"`
	DATABASES  []string  `json:"databases"`
	SEX        string    `json:"sex"`
	BIRTHDATE  string    `json:"birthdate"`
	ADDRESS    string    `json:"address"`
	ZIPCODE    string    `json:"zipcode"`
	CITY       string    `json:"city"`
	STATE      string    `json:"state"`
	COUNTRY    string    `json:"country"`
	CREATEDAT  time.Time `json:"created_at"`
	UPDATEDAT  time.Time `json:"updated_at"`
}
