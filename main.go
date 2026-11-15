package main

import (
	"log"

	"restapi/server"
)

func main() {
	srv := server.New(":5000", ":5001", "certs/cert.pem", "certs/key.pem")
	log.Fatal(srv.Start())
}
