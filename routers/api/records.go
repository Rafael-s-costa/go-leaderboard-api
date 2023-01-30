package api

import (
	"go-leaderboard-api/controllers"
	"go-leaderboard-api/pkg/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGameRecords(ctx *gin.Context) {
	var request requests.GetGameRecordsRequest

	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, controllers.GetGameRecords(&request))
		return
	}

	ctx.IndentedJSON(http.StatusOK, controllers.GetGameRecords(&request))
}

func RegisterRecord(ctx *gin.Context) {
	// TODO: Implement
}

func ConvertJsonToRequestStruct() {
	// TODO: Implement
}
