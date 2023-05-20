package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	DispayName sql.NullString
}
