package models

type RoomDetail struct {
	ID        int     `json:"id"`
	Account   Account `json:"account"`
	Game      Game    `json:"game"`
	Room_Name string  `json:"name"`
}

type RoomDetailsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type RoomsDetailsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Room   `json:"data"`
}