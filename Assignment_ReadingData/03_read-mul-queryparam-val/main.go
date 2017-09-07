package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

type dogs struct {
	Color string `json:"color"`
	Breed string `json:"breed"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dogs", ReadMulQueryParam).Methods("GET") //read multiple query parameters
	log.Fatal(http.ListenAndServe(":8087", router))
}

func ReadMulQueryParam(res http.ResponseWriter, req *http.Request) {
	fmt.Println("reading multiple Query parameter")
	var dog dogs
	fmt.Println("Raw Query", req.URL.RawQuery)
	QueryParam, err := url.ParseQuery(req.URL.RawQuery)
	fmt.Println(QueryParam)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	dog.Color = QueryParam["color"][0]
	dog.Breed = QueryParam["breed"][0]
	fmt.Fprintln(res, "Color=", dog.Color)
	fmt.Fprintln(res, "Breed=", dog.Breed)

}
