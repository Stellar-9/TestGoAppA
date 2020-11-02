package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//declaring test global variable

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

	keys, ok := r.URL.Query()["name"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]

	log.Println("Url Param 'key' is: " + string(key))
	w.Header().Set("Content-Type", "application/json") // setup JSON header
	json.NewEncoder(w).Encode("Hello " + string(key))  // encoding string and key value to JSON

}
