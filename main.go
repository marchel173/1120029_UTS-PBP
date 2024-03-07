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
	router.HandleFunc("/users", controllers.InsertUser).Methods("POST")
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc("/users/gorm", controllers.InsertUserGorm).Methods("POST")
	router.HandleFunc("/users/gorm", controllers.GetAllUsersGorm).Methods("GET")
	router.HandleFunc("/users/gorm/{id}", controllers.UpdateUserGorm).Methods("PUT")
	router.HandleFunc("/users/gorm/{id}", controllers.DeleteUserGorm).Methods("DELETE")
	
	router.HandleFunc("/products", controllers.InsertProduct).Methods("POST")
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	
	router.HandleFunc("/transactions", controllers.InsertTransaction).Methods("POST")
	router.HandleFunc("/transactions", controllers.GetAllTransactions).Methods("GET")
	router.HandleFunc("transactions/user", controllers.GetDetailUserTransaction).Methods("GET")
	router.HandleFunc("/transactions/{id}", controllers.UpdateTransaction).Methods("PUT")
	router.HandleFunc("/transactions/{id}", controllers.DeleteTransaction).Methods("DELETE")
	
	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}