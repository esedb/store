package main

import (
	"estore/common"
	"estore/routers"
	f "fmt"
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
	f.Println("server listening on: ", server.Addr)
	server.ListenAndServe()
	// Live Reloading gin --appPort 7000 -p 3000

}
