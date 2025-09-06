package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"uniqueIndex" json:"username"`
	Password   string `json:"password"`
	Email      string `gorm:"uniqueIndex" json:"email"`
	CreatedAt  int64  `json:"created_at" gorm:"autoCreateTime"`
	ModifiedAt int64  `json:"modified_at" gorm:"autoUpdateTime"`
}
