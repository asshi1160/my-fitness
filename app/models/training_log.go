package models

import (
	"time"
)

type TrainingLog struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint `gorm:"not null"`
	User       User
	TrainingID uint `gorm:"not null"`
	Training   Training
	RecordedAt time.Time `gorm:"not null;default=time.Now()"`
	Weight     float32
	Rep        int
	Set        int
	Distance   float32
	Time       time.Time
	Memo       string `gorm:"type:varchar(100)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
