package controllers

import (
	"go-leaderboard-api/pkg/errors"
	"go-leaderboard-api/pkg/requests"
	"go-leaderboard-api/pkg/responses"
)

func GetGameRecords(getGameRecordsRequest *requests.GetGameRecordsRequest) responses.BaseGetRecordsResponse {
	var response responses.BaseGetRecordsResponse
	response.ErrorCode = errors.SUCCESS

	return response
}
