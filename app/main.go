package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-fitness/database/connection"
	"my-fitness/models"
)

func main() {
	connection.Open()
	defer connection.Close()

	db := connection.GetConn()
	err := db.AutoMigrate(
		&models.User{},
		&models.Weight{},
		&models.Training{},
		&models.TrainingLog{},
	)
	if err != nil {
		panic("migration error")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run()
}
