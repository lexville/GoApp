package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalln("Failed to start the server")
	}
}
