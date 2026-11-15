package main

import (
	"log"

	_ "restapi/docs"
	"restapi/server"
)

//go:generate swag init

//	@title			REST API
//	@version		1.0
//	@description	Basic Go REST API.
//	@schemes		http https
//	@BasePath		/
func main() {
	srv := server.New(":5000", ":5001", "certs/cert.pem", "certs/key.pem")
	log.Fatal(srv.Start())
}
