package models

type Participant struct {
	ID        int `json:"id"`
	RoomID    int `json:"roomid"`
	AccountID int `json:"accountid"`
}

type ParticipantsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Participant `json:"data"`
}

type ParticipantResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Participant `json:"data"`
}
