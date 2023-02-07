package requests

type RegisterRecordRequest struct {
	GameID uint64 `json:"gameid" binding:"required"`
	UserID uint64 `json:"userid" binding:"required"`
	Score  uint64 `json:"score"`
	Time   uint64 `json:"time"`
}
