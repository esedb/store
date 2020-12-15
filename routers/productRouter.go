package routers

import (
	"estore/controllers"

	"github.com/gorilla/mux"
)

//UserRouter You should come back and attach authentication middleware
func ProductRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/product", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/search/{name}", controllers.SearchProductByName).Methods("GET")
	router.HandleFunc("/product/id/{product_id}", controllers.GetProduct).Methods("GET")
	router.HandleFunc("/product/{store_id}", controllers.GetAllProducts).Methods("GET")

	return router

}
