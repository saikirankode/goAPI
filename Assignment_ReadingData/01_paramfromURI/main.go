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
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleDogs(res http.ResponseWriter, req *http.Request) {

	fmt.Println("reading parameter from URI !")

	vars := mux.Vars(req)
	idkey := vars["idkey"]

	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, "dog param:", idkey)

}
