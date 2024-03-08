package models

type Game struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max_player"`
}

type GamesResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Game `json:"data"`
}

type GameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Game   `json:"data"`
}
