package controllers

import (
	"encoding/json"
	"github.com/foruscommunity/go-rest-api-example/pkg/db"
	"github.com/foruscommunity/go-rest-api-example/pkg/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var mockData = db.MockData()

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(mockData)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range mockData {
		if strconv.Itoa(item.Id) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Id = len(mockData) + 1
	mockData = append(mockData, user)
	json.NewEncoder(w).Encode(user)
}

func EditUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range mockData {
		if strconv.Itoa(item.Id) == params["id"] {
			var user models.User
			user.Id = item.Id
			_ = json.NewDecoder(r.Body).Decode(&user)
			mockData[index].Name = user.Name
			mockData[index].Username = user.Username
			mockData[index].Bio = user.Bio
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(mockData)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range mockData {
		if strconv.Itoa(item.Id) == params["id"] {
			mockData = append(mockData[:index], mockData[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(mockData)
}
