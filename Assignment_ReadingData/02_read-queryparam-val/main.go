package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

type dogs struct {
	ID    string `json:"dogID"`
	Color string `json:"color"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dogs/{dogID}", ReadURIParameter).Methods("GET")
	router.HandleFunc("/dogs", ReadQueryParameter).Methods("GET")
	log.Fatal(http.ListenAndServe(":8084", router))
}

func ReadURIParameter(res http.ResponseWriter, req *http.Request) {

	fmt.Println("reading URI parameter")
	var dog dogs
	param := mux.Vars(req)
	dog.ID = param["dogID"]
	json.NewEncoder(res).Encode(dog.ID)

}

func ReadQueryParameter(res http.ResponseWriter, req *http.Request) {
	var dog dogs
	fmt.Println("Raw Query", req.URL.RawQuery) //Query parses RawQuery and returns the corresponding values
	QueryParam, err := url.ParseQuery(req.URL.RawQuery)
	//ParseQuery parses the URL-encoded query string and returns a map listing the values specified for each key.
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	dog.Color = QueryParam["color"][0]
	fmt.Fprintln(res, "Dogcolor=", dog.Color)

}
