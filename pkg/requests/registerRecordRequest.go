package requests

type RegisterRecordRequest struct {
	GameID uint64 `json:"gameid"`
	UserID uint64 `json:"userid"`
	Score  uint64 `json:"score"`
	Time   uint64 `json:"time"`
}
