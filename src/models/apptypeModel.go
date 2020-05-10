package models

import (
	"time"
)

type Apptype struct {
	ID          string    `json:"id"`
	NAME        string    `json:"name"`
	DESCRIPTION string    `json:"description"`
	AUTHOR      string    `json:"author"`
	CREATEDAT   time.Time `json:"created_at"`
	UPDATEDAT   time.Time `json:"updated_at"`
}
