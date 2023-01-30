package requests

type GetGameUserRecordsRequest struct {
	GameID uint64 `json:"gameid"`
	UserID uint64 `json:"userid"`
}
