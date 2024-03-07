package controllers

import (
	"UTS/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func LoginUserGorm(w http.ResponseWriter, r *http.Request) {
	db:=connectGorm()
	var response models.ErrorResponse

	err := r.ParseForm()
	if err != nil {
		response.Status = 400
		response.Message = "Error Parsing Data"
		w.WriteHeader(400)
		log.Println(err.Error())
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	if email == "" || password == "" {
		response.Status = 400
		response.Message = "Please Input Email and Password"
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		response.Status = 400
		response.Message = "User not found"
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if user.Password == password {
		response.Status = 200
		response.Message = "Login Success"
		w.Header().Set("Content-Type", "application/json")
	} else {
		response.Status = 400
		response.Message = "Login Failed"
		w.WriteHeader(400)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllUsersGorm(w http.ResponseWriter, r *http.Request){
	db:=connectGorm()
	var users []models.User

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	query := db
	if name != "" {
		query = query.Where("name = ?", name)
	}
	if age != "" {
		ageInt, _ := strconv.Atoi(age)
		query = query.Where("age = ?", ageInt)
	}

	if err := query.Find(&users).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Internal Server Error"})
		return
	}

	w.Header().Set("Content-Type",  "application/json")
	json.NewEncoder(w).Encode(models.UsersResponse{Status: http.StatusOK, Message: "Successfully Retrieved Users Data.", Data: users})
}

func DeleteUserGorm(w http.ResponseWriter, r *http.Request) {
	db:=connectGorm()
	vars := mux.Vars(r)
	userId := vars["id"]

	var user models.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Status: http.StatusBadRequest, Message: "User not found"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Failed Delete Data"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ErrorResponse{Status: http.StatusOK, Message: "Success Delete Data"})
}

func InsertUserGorm(w http.ResponseWriter, r *http.Request) {
	db:=connectGorm()
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Status: http.StatusBadRequest, Message: "Error Parsing Data"})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorResponse{Status: http.StatusInternalServerError, Message: "Error Insert Data"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserResponse{Status: http.StatusOK, Message: "Success", Data: user})
}

func UpdateUserGorm(w http.ResponseWriter, r *http.Request) {
	db:=connectGorm()
	vars := mux.Vars(r)
	userId := vars["id"]

	var user models.User
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Status: http.StatusBadRequest, Message: "User not found"})
		return
	}

	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{Status: http.StatusBadRequest, Message: "Error Parsing Data"})
		return
	}

	db.Model(&user).Updates(newUser)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.UserResponse{Status: http.StatusOK, Message: "Success Update Data", Data: newUser})
}
