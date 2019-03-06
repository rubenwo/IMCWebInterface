package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/config/update", updateConfigEndpoint).Methods("POST")
	router.HandleFunc("/config/delete{id}", deleteConfigEndpoint).Methods("DELETE")
	router.HandleFunc("/config/get{id}", getConfigEndpoint).Methods("GET")

	fmt.Println("Starting IMC server on port 80...")
	log.Fatal(http.ListenAndServe(":80", router))
}

func getConfigEndpoint(w http.ResponseWriter, req *http.Request) {

}
func deleteConfigEndpoint(w http.ResponseWriter, req *http.Request) {

}
func updateConfigEndpoint(w http.ResponseWriter, req *http.Request) {

}
