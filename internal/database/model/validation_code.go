package model

import "gorm.io/gorm"

type ValidationCode struct {
	gorm.Model
	Code   string `gorm:"size:20;not null"`
	Email  string `gorm:"size:255;not null"`
	UsedAt int
}
