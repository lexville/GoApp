package main

import (
	"GoApp/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	homeController := controllers.AddViewTemplate()
	r.HandleFunc("/", homeController.Home).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalln("Failed to start the server")
	}
}
