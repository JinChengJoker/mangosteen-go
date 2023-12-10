package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"size:255;not null"`
}
