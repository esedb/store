package routers

import (
	"estore/controllers"

	"github.com/gorilla/mux"
)

//UserRouter You should come back and attach authentication middleware
func CategoryRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/category", controllers.CreateCategory).Methods("POST")
	router.HandleFunc("/category", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{name}", controllers.SearchCategoryByName).Methods("GET")

	return router

}
