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

type storeController struct{}

var StoreController = storeController{}

func (controller storeController) CreateStore(w http.ResponseWriter, r *http.Request) {
	var store model.Store
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	er := service.StoreService.CreateStore(&store)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to create a Proudct at this moment, please try later",
			500)

		return

	}
	resp, err := json.Marshal(&store)
	if err != nil {
		log.Fatalf("[CreateUserController] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func (controller storeController) UpdateStore(w http.ResponseWriter, r *http.Request) {
	var store model.Store
	err := json.NewDecoder(r.Body).Decode(&store)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	er := service.StoreService.UpdateStore(&store)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to update a Product at this moment, please try later",
			500)

		return

	}
	resp, err := json.Marshal(&store)
	if err != nil {
		log.Fatalf("[UpdateProductcontroller] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func (controller storeController) GetAllStore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchant_id := vars["merchant_id"]

	stores, err := service.StoreService.GetAllStore(merchant_id)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"An error occured.",
			500)

		return

	}
	resp, err := json.Marshal(&stores)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"An error occured when trying to encode data.",
			500)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}
