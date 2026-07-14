package models

import "time"

type User struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
