package models

import (
	"time"

	"gorm.io/gorm"
)

type Habit struct {
	gorm.Model
	Name        string    `json:"title"`
	Description string    `json:"content"`
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
