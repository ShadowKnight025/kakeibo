package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"

	"lionheart.dev/entity"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := router.Run(); err != nil{
		log.Fatal("Failed to start server: %v", err)
	}
}
