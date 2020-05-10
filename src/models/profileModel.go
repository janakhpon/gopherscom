package models

import (
	"time"
)

type Profile struct {
	ID         string    `json:"id"`
	USERID     string    `json:"userid"`
	CAREER     string    `json:"career"`
	FRAMEWORKS []string  `json:"frameworks"`
	LANGUAGES  []string  `json:"languages"`
	PLATFORMS  []string  `json:"platforms"`
	DATABASES  []string  `json:"databases"`
	OTHERS     []string  `json:"others"`
	SEX        string    `json:"sex"`
	BIRTHDATE  string    `json:"birthdate"`
	ADDRESS    string    `json:"address"`
	ZIPCODE    string    `json:"zipcode"`
	CITY       string    `json:"city"`
	STATE      string    `json:"state"`
	COUNTRY    string    `json:"country"`
	LAT        float64   `json:"lat"`
	LON        float64   `json:"lon"`
	CREATEDAT  time.Time `json:"created_at"`
	UPDATEDAT  time.Time `json:"updated_at"`
}
