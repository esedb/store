package main

import (
	"estore/common"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	common.StartUp()
	n := negroni.Classic()
	server := http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	server.ListenAndServe()

}
