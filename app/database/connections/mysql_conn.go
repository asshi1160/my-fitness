package connections

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type MySQL struct{}

func (my MySQL) Conn() (*sql.DB, error) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	conf := mysql.Config{
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      "mysql",
		DBName:    os.Getenv("DB_NAME"),
		Loc:       jst,
		ParseTime: true,
		Collation: "utf8mb4_bin",
	}

	db, _ := sql.Open("mysql", conf.FormatDSN())
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err := db.PingContext(ctx)
	if err != nil {
		log.Printf("error: %s", err)
	}

	return db, err
}
