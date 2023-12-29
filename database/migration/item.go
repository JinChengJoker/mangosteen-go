package migration

import (
	"mangosteen/dal/model"
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	UserId int64 `gorm:"not null"`
	User   model.User
	PaidAt *time.Time
}
