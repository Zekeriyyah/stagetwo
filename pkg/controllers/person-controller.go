package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"log"

	"github.com/Zekeriyyah/stagetwo/pkg/models"
	"github.com/Zekeriyyah/stagetwo/pkg/utils"
	"github.com/gorilla/mux"
)

var NewUser models.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	newPersons := models.GetAllUsers()
	err := json.NewEncoder(w).Encode(&newPersons)
	if err != nil {
		fmt.Println("Error while Encoding!")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userId := params["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	
	userDetails, db, err := models.GetUserById(ID)
	log.Println("Test Error: ", err)
	if err != nil {
		fmt.Fprintf(w, "Sorry, User with the id %v is not found!", ID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(userDetails)
	if err != nil {
		log.Println("Error while encoding response")
		return
	}
	w.WriteHeader(http.StatusOK)
	
	
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	createUser := &models.User{}
	utils.ParseBody(r, createUser)
	b := createUser.CreateUser()
	err := json.NewEncoder(w).Encode(&b)
	if err != nil {
		log.Println("Error: Encoding failed")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	userId := params["userId"]
	ID, _ := strconv.ParseInt(userId, 0, 0)

	deletedUserErr := models.DeleteUser(ID)
	if deletedUserErr == nil {
		res := "User Deleted Successfully"
		resbyte, _ := json.Marshal(res)
		w.Write(resbyte)
	} else {
		res := "Failed to delete User: "
		resbyte, _ := json.Marshal(res)
		w.Write(resbyte)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateUser = &models.User{}
	utils.ParseBody(r, updateUser)
	params := mux.Vars(r)
	userId := params["userId"]

	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("Error: Error while parsing")
	}

	userDetails, db, err := models.GetUserById(ID)
	if err != nil {
		fmt.Fprintf(w, "User with Id %v not found", ID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userDetails.Name != "" {
		userDetails.Name = updateUser.Name
	}
	if userDetails.Email != "" {
		userDetails.Email = updateUser.Email
	}
	if userDetails.Country != "" {
		userDetails.Country = updateUser.Country
	}
	db.Save(&userDetails)
	_ = json.NewEncoder(w).Encode(&userDetails)
	w.WriteHeader(http.StatusOK)
}
