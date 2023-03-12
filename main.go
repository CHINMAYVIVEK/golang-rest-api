package main

import (
	server "go-docker/server"

	"github.com/gorilla/mux"
)

func main() {

	// =================== DB Connection =========
	// fmt.Printf("%+v", jdconfig.Cfg)
	// jdconfig.Cfg.DbConnect()
	// defer jdconfig.CloseDbJds()
	// defer jdconfig.CloseDbJustdial()

	// =================== DB Connection =========

	r := mux.NewRouter()

	server.InitializeRouter(r)

	server.ServerStart(r)

}
