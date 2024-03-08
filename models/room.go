package models

type Room struct {
	ID        int    `json:"id"`
	AccountID int    `json:"account_id"`
	GameID    int    `json:"gameid"`
	Room_Name string `json:"room_name"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}