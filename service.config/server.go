package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Configs struct {
	Configs []Config `json:"configs"`
}

var configs Configs

func main() {
	loadConfigs()
	router := mux.NewRouter()

	router.HandleFunc("/config/update/{id}", updateConfigEndpoint).Methods("POST")
	router.HandleFunc("/config/delete/{id}", deleteConfigEndpoint).Methods("DELETE")
	router.HandleFunc("/config/get/{id}", getConfigEndpoint).Methods("GET")

	fmt.Println("Starting IMC server on port 80...")
	log.Fatal(http.ListenAndServe(":80", router))
}

func loadConfigs() {
	fmt.Println("Loading configs...")
	jsonFile, err := os.Open("configs.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened configs.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &configs)
	fmt.Printf("Loaded %d configs.\n", len(configs.Configs))
}

func getConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Printf("Got request for config with ID: %s.\n", id)
}

func deleteConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Printf("Got request for config with ID: %s.\n", id)
}

func updateConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Printf("Got request for config with ID: %s.\n", id)
}
