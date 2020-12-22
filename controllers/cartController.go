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

func AddToCart(w http.ResponseWriter, r *http.Request) {
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

	er := cartService.AddToCart(cart)
	if err != nil {
		common.DisplayError(
			w,
			er,
			"Unable to add Items to Cart, please try again later or contact the administrator.",
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
