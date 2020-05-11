package models

import (
	"time"
)

type Bootcamp struct {
	ID          string     `json:"id"`
	TOPIC       string     `json:"topic"`
	INSTRUCTORS []string   `json:"instructors"`
	ADDRESS     string     `json:"address"`
	LAT         float64    `json:"lat"`
	LON         float64    `json:"lon"`
	STUDENTS    string     `json:"students"`
	ENROLLMENTS []Enroller `json:"enrollments"`
	DESCRIPTION string     `json:"description"`
	AVAILABLE   bool       `json:"available"`
	STARTEDAT   string     `json:"startedat"`
	FINISHEDAT  string     `json:"finishedat"`
	AUTHOR      string     `json:"author"`
	LIKES       []Like     `json:"likes"`
	COMMENTS    []Comment  `json:"comments"`
	CREATEDAT   time.Time  `json:"created_at"`
	UPDATEDAT   time.Time  `json:"updated_at"`
}

type Enroller struct {
	ID   string `json:"id"`
	NAME string `json:"name"`
}

type Like struct {
	ID   string `json:"id"`
	NAME string `json:"name"`
	LIKE bool   `json:"like"`
}

type Comment struct {
	ID        string    `json:"id"`
	NAME      string    `json:"name"`
	TEXT      string    `json:"text"`
	EDITED    bool      `json:"edited"`
	UPDATEDAT time.Time `json:"updated_at"`
}
