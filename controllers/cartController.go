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
	"gopkg.in/go-playground/validator.v9"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
	var cartService = &service.CartService{}
	fmt.Println("after here")
	var cart model.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	_validator := validator.New()
	verr := _validator.Struct(cart)
	if verr != nil {
		common.DisplayError(
			w,
			verr,
			"Validation Error!",
			400)

		return

	}

	result, er := cartService.AddToCart(cart)
	if er != nil {
		common.DisplayError(
			w,
			er,
			"An Error occured trying to add Cart Items",
			500)

		return

	}
	if result != "OK" {
		common.DisplayError(
			w,
			er,
			"Adding Items to Cart Failed!",
			500)

		return

	}
	resp, err := json.Marshal(&cart)
	if err != nil {
		log.Fatalf("[AddToCart] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func GetItemsFromCart(w http.ResponseWriter, r *http.Request) {
	var cartService = &service.CartService{}
	vars := mux.Vars(r)
	user_id := vars["user_id"]
	cart, err := cartService.GetFromCart(user_id)
	resp, err := json.Marshal(&cart)
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

func RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	var cartService = &service.CartService{}
	var cart model.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		common.DisplayError(
			w,
			err,
			"Can't able to Deserialize data type, make sure your data input is correct",
			500)

		return
	}
	_validator := validator.New()
	verr := _validator.Struct(cart)
	if verr != nil {
		common.DisplayError(
			w,
			verr,
			"Validation Error!",
			400)

		return

	}

	result, er := cartService.RemoveFromCart(cart)
	if er != nil {
		common.DisplayError(
			w,
			er,
			"An Error occured trying to add Cart Items",
			500)

		return

	}
	if result != "OK" {
		common.DisplayError(
			w,
			er,
			"Adding Items to Cart Failed!",
			500)

		return

	}
	resp, err := json.Marshal(&cart)
	if err != nil {
		log.Fatalf("[RemoveFromCart] %s\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}
