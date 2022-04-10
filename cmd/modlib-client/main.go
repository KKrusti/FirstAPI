package main

import (
	"apirest/cmd/modlib-server"
	"log"
	"net/http"
)

func main() {
	s := modlib_server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
