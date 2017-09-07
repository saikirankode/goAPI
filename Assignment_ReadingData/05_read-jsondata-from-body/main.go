package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"net/url"
)

type dogs struct {
	Color string `json:"color"`
	Breed string `json:"breed"`
}

func main() {

	http.HandleFunc("/dogs", QueryParamWithMulVal) //read query parameters with multiple values
	http.HandleFunc("/readdogs", ReadDataFrmBody)
	//read the data from body (json data) and store it in a structure variables
	http.HandleFunc("/readFrmHeader", ReadDataFrmHeader) //read value in a header and print it.

	log.Fatal(http.ListenAndServe(":8100", nil))

}

func QueryParamWithMulVal(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Raw Query", req.URL.RawQuery)
	var dog dogs
	QueryParam, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	dog.Color = QueryParam["color"][0]
	//dog.Color = QueryParam["color"][1]

	fmt.Fprintln(res, "Color=", dog.Color)

}

func ReadDataFrmBody(res http.ResponseWriter, req *http.Request) {
	var dog dogs
	params := req.Body
	json.NewDecoder(params).Decode(&dog)
	fmt.Fprintln(res, "DogColor=", dog.Color)
	fmt.Fprintln(res, "DogBreed=", dog.Breed)
}

func ReadDataFrmHeader(res http.ResponseWriter, req *http.Request) {
	fmt.Println("reading from header")
	var dog dogs
	color := req.Header.Get("Color")
	breed := req.Header.Get("Breed")
	dog.Color = string(color)
	dog.Breed = string(breed)
	fmt.Fprintln(res, "DogColor=", dog.Color)
	fmt.Fprintln(res, "DogBreed=", dog.Breed)

}
