package routers

import (
	"estore/controllers"

	"github.com/gorilla/mux"
)

//UserRouter You should come back and attach authentication middleware
func ItemRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/item", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/item", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/item/search/{name}", controllers.SearchItemByName).Methods("GET")
	router.HandleFunc("/item/id/{item_id}", controllers.GetItem).Methods("GET")
	router.HandleFunc("/item/{store_id}", controllers.GetAllItemsByStore).Methods("GET")
	router.HandleFunc("/item/{store_id}", controllers.GetAllItems).Methods("GET")

	return router

}
