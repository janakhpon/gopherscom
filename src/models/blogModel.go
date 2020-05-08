package models

import (
	"time"
)

type Blog struct {
	ID        string    `json:"id"`
	TITLE     string    `json:"title"`
	BODY      string    `json:"body"`
	PUBLIC    bool      `json:"public"`
	APPTYPE   string    `json:"apptype"`
	LANGUAGES []string  `json:"languages"`
	TAGS      []string  `json:"tags"`
	LIBRARIES []string  `json:"libraries"`
	AUTHOR    string    `json:"author"`
	CREATEDAT time.Time `json:"created_at"`
	UPDATEDAT time.Time `json:"updated_at"`
}
