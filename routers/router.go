package routers

import (
	"github.com/gorilla/mux"
)

//UserRouter You should come back and attach authentication middleware
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = UserRouter(router)
	router = CategoryRouter(router)
	router = ItemRouter(router)
	router = StoreRouter(router)
	router = CartRouter(router)

	return router

}
