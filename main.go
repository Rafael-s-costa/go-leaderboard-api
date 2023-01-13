package main

// album represents data about a record album.
type user struct {
	ID    string `json:"id"`
	name  string `json:"name"`
	email string `json:"email"`
}

type game struct {
	ID   string `json:"id"`
	name string `json:"name"`
}

type record struct {
	ID     string `json:"id"`
	time   string `json:"time"`
	score  string `json:"score"`
	gameID string `json:"gameid"`
	userID string `json:"userid"`
}
