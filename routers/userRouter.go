package routers

import (
	"estore/controllers"

	"github.com/gorilla/mux"
)

//UserRouter You should come back and attach authentication middleware
func UserRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{email}", controllers.GetUserByEmail).Methods("GET")

	return router

}
