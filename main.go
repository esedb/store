package main

import (
	"estore/common"
	"estore/routers"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	common.StartUp()
	n := negroni.Classic()
	router := routers.InitRoutes()
	n.UseHandler(router)
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	server.ListenAndServe()

}
