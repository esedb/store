package controllers

import (
	"encoding/json"
	"estore/common"
	"estore/model"
	"estore/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	_validator := validator.New()
	verr := _validator.Struct(item)
	if verr != nil {
		common.DisplayError(
			w,
			verr,
			"Validation Error!",
			400)

		return

	}

	er := service.ItemService.CreateItem(&item)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to create a Proudct at this moment, please try later",
			500)

		return

	}
	resp, err := json.Marshal(&item)
	if err != nil {
		log.Fatalf("[CreateUserController] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}

	_validator := validator.New()
	verr := _validator.Struct(item)
	if verr != nil {
		common.DisplayError(
			w,
			verr,
			"Validation Error!",
			400)

		return

	}

	er := service.ItemService.UpdateItem(&item)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to update a Item at this moment, please try later",
			400)

		return

	}
	resp, err := json.Marshal(&item)
	if err != nil {
		log.Fatalf("[UpdateItemcontroller] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func SearchItemByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	item, err := service.ItemService.SearchItemByName(name)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"An error occured.",
			500)

		return
	}
	resp, err := json.Marshal(&item)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func GetAllItemsByStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	store_id := vars["store_id"]
	items := service.ItemService.GetAllItemsByStore(store_id)
	resp, err := json.Marshal(&items)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"An error occured.",
			500)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	item_id := vars["item_id"]
	item, err := service.ItemService.GetItem(item_id)
	resp, err := json.Marshal(&item)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"An error occured.",
			500)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := service.ItemService.GetAllItems()
	if err != nil {
		common.DisplayError(
			w,
			err,
			"An error occured.",
			500)

		return

	}

	resp, _ := json.Marshal(&items)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}
