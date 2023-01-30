package responses

import "go-leaderboard-api/models"

type BaseGetRecordsResponse struct {
	ErrorCode uint32           `json:"errorcode"`
	Records   *[]models.Record `json:"records"`
}
