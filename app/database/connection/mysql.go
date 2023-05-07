package connection

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Open() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(mysql:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
}

func Close() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func GetConn() *gorm.DB {
	return db
}
