package models

import (
	"time"
)

type Weight struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	User       User
	Weight     float32 `gorm:"not null"`
	BodyFat    *float32
	RecordedAt time.Time
	CreatedAt  time.Time
}
