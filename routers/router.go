package routers

import (
	"log"

	"go-leaderboard-api/routers/api"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {
	log.Println("Initializing routes")
	router := gin.Default()

	router.GET("/records", api.GetGameRecords)
	router.POST("/records", api.RegisterRecord)

	router.Run()
}
