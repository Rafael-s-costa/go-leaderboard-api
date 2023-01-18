package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGameRecords(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "hhhh"})
}
