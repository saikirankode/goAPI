package main

import (
	//"encoding/json"
	"fmt"

	"log"
	"net/http"
	//"net/url"
	//"encoding/json"
	"io/ioutil"
)

/*type dogs struct {
	Color string `json:"color"`
	Breed string `json:"breed"`
}*/
//var dog string
func main() {

	//http.HandleFunc("/dogs", QueryParamWithMulVal)
	http.HandleFunc("/readdogs", ReadDataFrmBody)

	log.Fatal(http.ListenAndServe(":8099", nil))

}

/*func QueryParamWithMulVal(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Raw Query", req.URL.RawQuery)

	QueryParam, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	dog.Color = QueryParam["color"][0]
	//dog.Color = QueryParam["color"][1]

	fmt.Fprintln(res, "Color=", dog.Color)

}*/
func ReadDataFrmBody(res http.ResponseWriter, req *http.Request) {
	fmt.Println("reading from body")

	dogcolor := req.Body

	d, err := ioutil.ReadAll(dogcolor)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	e := string(d)
	fmt.Fprintln(res, "DogColor=", e)

}
