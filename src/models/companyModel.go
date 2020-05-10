package models

import (
	"time"
)

type Company struct {
	ID          string    `json:"id"`
	NAME        string    `json:"name"`
	PRODUCTS    []string  `json:"products"`
	EMPLOYEE    string    `json:"employee"`
	BRANCHES    []*Branch `pg:",one2many:company_branches"`
	FRAMEWORKS  []string  `json:"frameworks"`
	LANGUAGES   []string  `json:"languages"`
	PLATFORMS   []string  `json:"platforms"`
	DATABASES   []string  `json:"databases"`
	OTHERS      []string  `json:"others"`
	ADDRESS     string    `json:"address"`
	ZIPCODE     string    `json:"zipcode"`
	CITY        string    `json:"city"`
	STATE       string    `json:"state"`
	COUNTRY     string    `json:"country"`
	LAT         float64   `json:"lat"`
	LON         float64   `json:"lon"`
	FOUNDEDYEAR string    `json:"foundedyear"`
	CREATEDAT   time.Time `json:"created_at"`
	UPDATEDAT   time.Time `json:"updated_at"`
}

type Branch struct {
	ID          string    `json:"id"`
	CID         string    `json:"cid"`
	NAME        string    `json:"name"`
	ADDRESS     string    `json:"address"`
	ZIPCODE     string    `json:"zipcode"`
	CITY        string    `json:"city"`
	STATE       string    `json:"state"`
	COUNTRY     string    `json:"country"`
	LAT         float64   `json:"lat"`
	LON         float64   `json:"lon"`
	FOUNDEDYEAR string    `json:"foundedyear"`
	CREATEDAT   time.Time `json:"created_at"`
	UPDATEDAT   time.Time `json:"updated_at"`
}
