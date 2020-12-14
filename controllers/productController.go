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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	er := service.ProductService.CreateProduct(&product)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to create a Proudct at this moment, please try later",
			500)

		return

	}
	resp, err := json.Marshal(&product)
	if err != nil {
		log.Fatalf("[CreateUserController] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	er := service.ProductService.UpdateProduct(&product)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to update a Product at this moment, please try later",
			500)

		return

	}
	resp, err := json.Marshal(&product)
	if err != nil {
		log.Fatalf("[UpdateProductcontroller] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func SearchProductByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	product, err := service.ProductService.SearchPRoductByName(name)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"An error occured.",
			500)

		return
	}
	resp, err := json.Marshal(&product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}
