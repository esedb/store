package routers

import (
	"estore/controllers"

	"github.com/gorilla/mux"
)

//UserRouter You should come back and attach authentication middleware
func CartRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/cart", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/cart", controllers.SearchItemByName).Methods("GET")

	return router

}
