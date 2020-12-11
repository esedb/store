package routers

import (
	"github.com/gorilla/mux"
)

//UserRouter You should come back and attach authentication middleware
func InitRoutes() *mux.Router {
	router := &mux.NewRouter()
	router = router.UserRouter(router)

	return router

}
