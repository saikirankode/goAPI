package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dogs/{idkey}", handleDogs).Methods("GET")
	http.ListenAndServe(":12345", router)
}

func handleDogs(res http.ResponseWriter, req *http.Request) {

	log.Println("reading parameter from URI")

	vars := mux.Vars(req)
	idkey := vars["idkey"]

	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, "dog param:", idkey)

}
