package models

import (
	"time"
)

type Training struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	User          User
	BodyPart      string   `gorm:"type:varchar(10);not null"`
	TrainingType  string   `gorm:"type:varchar(20);not null"`
	TrainingName  string   `gorm:"type:varchar(30);unique;not null"`
	TrainingForms []string `gorm:"type:varchar(50);not null;serializer:json"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
