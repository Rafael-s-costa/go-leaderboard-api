package models

type Record struct {
	ID     string `json:"id"`
	time   string `json:"time"`
	score  string `json:"score"`
	gameID string `json:"gameid"`
	userID string `json:"userid"`
}
