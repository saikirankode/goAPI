package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/dogs/{key}/{var1}", handleDogs).Methods("GET")
	http.ListenAndServe(":10000", router)
}

func handleDogs(res http.ResponseWriter, req *http.Request) {

	log.Println("reading parameter from URI")

	vars := mux.Vars(req)
	key := vars["key"]
	var1 := vars["var1"]

	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, "dog id :", key, var1)

}
