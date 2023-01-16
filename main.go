package main

import (
	"net/http"

	"go-leaderboard-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	var games = []models.Game{
		{1, "f"},
		{2, "test"},
	}

	router := gin.Default()
	router.GET("/games", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": games})
	})

	router.Run()
}
