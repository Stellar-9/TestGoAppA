package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	// get ENV variable value for os
	var listenPort string
	listenPort = os.Getenv("PORT")

	// create mux router
	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorld).Methods("GET") //use mux method for request
	http.ListenAndServe(":"+listenPort, r)       //create webserver

}

// create function for JSON format

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	names, ok := r.URL.Query()["name"]

	if !ok || len(names[0]) < 1 {
		log.Println("Missing URL parameter 'name'")
		return
	}

	name := names[0]

	log.Println("URL parameter 'name': " + string(name))
	w.Header().Set("Content-Type", "application/json") // setup JSON header
	json.NewEncoder(w).Encode("Hello " + string(name)) // encoding string and name value to JSON

}
