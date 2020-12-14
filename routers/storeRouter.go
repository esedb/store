package routers

import (
	"estore/controllers"

	"github.com/gorilla/mux"
)

//UserRouter You should come back and attach authentication middleware
func StoreRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/store", controllers.StoreController.CreateStore).Methods("POST")
	router.HandleFunc("/store", controllers.StoreController.UpdateStore).Methods("PUT")
	router.HandleFunc("/store", controllers.StoreController.GetAllStore).Methods("GET")

	return router

}
