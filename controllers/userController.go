package controllers

import (
	"encoding/json"
	"estore/common"
	"estore/model"
	"estore/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//CreateUser controller for createing User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	er := service.CreateUser(&user)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to create User at this moment, please try later",
			500)

		return

	}
	resp, err := json.Marshal(&user)
	if err != nil {
		log.Fatalf("[CreateUserController] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

//UpdateUser update uer method
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	er := service.UpdateUser(&user)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to create User at this moment, please try later",
			500)

		return

	}
	resp, err := json.Marshal(&user)
	if err != nil {
		log.Fatalf("[CreateUserController] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

//GetUserByEmail return user base on email (email is unique)
func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	user, er := service.GetUserByEmail(email)
	if er != nil {
		common.DisplayError(
			w,
			er,
			"An error occured retrieving this user",
			500)

		return

	}
	resp, err := json.Marshal(&user)
	if err != nil {
		log.Fatalf("[CreateUserController] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}
