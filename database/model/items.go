package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	UserId int64 `gorm:"not null"`
}
