package controllers

import (
	m "UTS/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()

	defer db.Close()
	var response m.RoomsResponse

	query := "SELECT * FROM rooms"
	id := r.URL.Query()["id"]
	if id != nil {
		query += " WHERE id = " + id[0]
	}

	rows, err := db.Query(query)

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var room m.Room
	var rooms []m.Room

	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.Room_Name, &room.GameID); err != nil {
			log.Println(err.Error())
		} else {
			rooms = append(rooms, room)
		}
	}

	if len(rooms) != 0 {
		response.Status = 200
		response.Message = "Succes Get Data"
		response.Data = rooms
	} else if response.Message == "" {
		response.Status = 400
		response.Message = "Data Not Found"
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	var response m.ErrorResponse

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	vars := mux.Vars(r)
	roomid := vars["id"]
	query, errQuery := db.Exec(`DELETE FROM rooms WHERE id = ?;`, roomid)
	RowsAffected, err := query.RowsAffected()

	if RowsAffected == 0 {
		response.Status = 400
		response.Message = "Room not found"
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success Delete Data"
		w.WriteHeader(200)
	} else {
		response.Status = 400
		response.Message = "Failed Delete Data"
		w.WriteHeader(400)
		log.Println(errQuery.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()
	var response m.RoomResponse
	err := r.ParseForm()

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	var room m.Room

	room.ID, _ = strconv.Atoi(r.Form.Get("id"))
	room.Room_Name = r.Form.Get("room_name")
	room.GameID, _ = strconv.Atoi(r.Form.Get("gameid"))

	rows, errQuery := db.Query("SELECT * FROM games WHERE id=?", room.GameID)

	if errQuery != nil {
		response.Status = 400
		response.Message = err.Error()
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	i := 0
	for rows.Next() {
		i++
	}

	if i == 0 {
		_, err = db.Exec("INSERT INTO games (id) VALUES (?)", room.GameID)

		if err != nil {
			response.Status = 400
			response.Message = err.Error()
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	res, errQuery := db.Exec("INSERT INTO room (room_name, id_game) VALUES (?,?,?)", room.ID, room.Room_Name, room.GameID)

	id, err := res.LastInsertId()

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
		room.ID = int(id)
		response.Data = room
	} else {
		response.Status = 400
		response.Message = "Error Insert Data"
		w.WriteHeader(400)
		log.Println(errQuery.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	var response m.RoomResponse

	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	vars := mux.Vars(r)
	roomid := vars["id"]

	var room m.Room
	room.ID, _ = strconv.Atoi(r.Form.Get("id"))
	room.Room_Name = r.Form.Get("room_name")
	room.GameID, _ = strconv.Atoi(r.Form.Get("gameid"))

	rows, _ := db.Query("SELECT * FROM rooms WHERE id = ?", roomid)
	var prevDatas []m.Room
	var prevData m.Room

	for rows.Next() {
		if err := rows.Scan(&prevData.ID, &prevData.Room_Name, &prevData.GameID); err != nil {
			log.Println(err.Error())
		} else {
			prevDatas = append(prevDatas, prevData)
		}
	}

	if len(prevDatas) > 0 {
		if room.GameID == 0 {
			room.GameID = prevDatas[0].GameID
		}

		_, errQuery := db.Exec(`UPDATE rooms SET Room_Name = ?, id_games = ? WHERE id = ?;`, room.Room_Name, room.GameID, roomid)

		if errQuery == nil {
			response.Status = 200
			response.Message = "Success Update Data"
			id, _ := strconv.Atoi(roomid)
			room.ID = id
			response.Data = room
			w.WriteHeader(200)
		} else {
			response.Status = 400
			response.Message = "Error Update Data"
			w.WriteHeader(400)
			log.Println(errQuery)

		}
	} else {
		response.Status = 400
		response.Message = "Data Not Found"
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetDetailAccountRoom(w http.ResponseWriter, r *http.Request) {

	db := connect()
	defer db.Close()

	var roomDetails []m.RoomDetail

	query := "SELECT r.ID , a.ID, a.Name, u.Age, u.Address, p.ID, p.Name, p.Price, t.Quantity FROM transactions t JOIN users u ON t.UserId = u.ID JOIN products p ON p.ID = t.ProductID"

	id := r.URL.Query()["id"]
	if id != nil {
		query += " WHERE u.id = " + id[0]
	}

	rows, err := db.Query(query)

	var response m.RoomDetailsResponse

	if err != nil {
		response.Status = 400
		response.Message = err.Error()
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var roomDetail m.RoomDetail
	var Account m.Account
	var Game m.Game

	for rows.Next() {
		if err := rows.Scan(&m.RoomDetail.ID, &Account.ID, &Account.Username, &Game.ID, &Game.Name, &Game.Max_player); err != nil {
			log.Println(err.Error())
		} else {
			roomDetail.Account = Account
			roomDetail.Game = Game
			roomDetails = append(roomDetails, roomDetail)
		}
	}

	if len(roomDetails) != 0 {
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = roomDetails
	} else {
		response.Status = 400
		response.Message = "Error Get Data"
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}