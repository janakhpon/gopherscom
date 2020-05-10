package models

import (
	"time"
)

type Library struct {
	ID          string    `json:"id"`
	NAME        string    `json:"name"`
	DESCRIPTION string    `json:"description"`
	LANGUAGES   []string  `json:"languages"`
	AUTHOR      string    `json:"author"`
	CREATEDAT   time.Time `json:"created_at"`
	UPDATEDAT   time.Time `json:"updated_at"`
}
