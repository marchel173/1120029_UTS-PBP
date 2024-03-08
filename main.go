package main

import (
	"UTS/controllers"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	
	router.HandleFunc("/rooms", controllers.InsertRoom).Methods("POST")
	router.HandleFunc("/rooms", controllers.GetAllRooms).Methods("GET")
	router.HandleFunc("rooms/account", controllers.GetDetailAccountRoom).Methods("GET")
	router.HandleFunc("/rooms/{id}", controllers.UpdateRoom).Methods("PUT")
	router.HandleFunc("/rooms/{id}", controllers.LeaveRoom).Methods("DELETE")
	
	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}