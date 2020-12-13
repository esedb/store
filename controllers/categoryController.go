package controllers

import (
	"encoding/json"
	"estore/common"
	"estore/model"
	"estore/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	er := service.CategoryService.CreateCategory(&category)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to create User at this moment, please try later",
			500)

		return

	}
	resp, err := json.Marshal(&category)
	if err != nil {
		log.Fatalf("[CreateUserController] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	fmt.Println("Body: ", r.Body.)
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	er := service.CategoryService.UpdateCategory(&category)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to create User at this moment, please try later",
			500)

		return

	}
	resp, err := json.Marshal(&category)
	if err != nil {
		log.Fatalf("[CreateUserController] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func SearchCategoryByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	category, err := service.CategoryService.SearchCategoryByName(name)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"An error occured.",
			500)

		return
	}
	resp, err := json.Marshal(&category)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}
