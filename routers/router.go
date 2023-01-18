package routers

import (
	"log"

	"go-leaderboard-api/routers/api"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {
	log.Println("Initializing routes")
	router := gin.Default()

	router.GET("/records/:gameId", api.GetGameRecords)

	router.Run()
}
