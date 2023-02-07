package requests

type GetGameUserRecordsRequest struct {
	GameID uint64 `json:"gameid" binding:"required"`
	UserID uint64 `json:"userid" binding:"required"`
}
