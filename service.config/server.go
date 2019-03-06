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

	router.HandleFunc("/config/{id}", updateConfigEndpoint).Methods("POST")
	router.HandleFunc("/config/{id}", deleteConfigEndpoint).Methods("DELETE")
	router.HandleFunc("/config/{id}", getConfigEndpoint).Methods("GET")

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

// Enables CORS for http responses
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// POST: Updates the config of 1 radio. Has a radioID as var. Receives JSON Config object as body.
func updateConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	// Get radioID from vars
	vars := mux.Vars(req)
	fmt.Println("Radio ID:", vars["radioID"])

	// Decode Config from body
	var c Config
	if err := json.NewDecoder(req.Body).Decode(&c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	// Begin Response
	enableCORS(&w)
	w.WriteHeader(200)
}

// GET: Get the config of 1 radio. Has a radioID as var.
func getConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	// Get radioID from vars
	vars := mux.Vars(req)
	fmt.Println("ID:", vars["id"])

	// Begin Response
	enableCORS(&w)
	w.WriteHeader(200)
}

// GET: Get the config of 1 radio. Has a radioID as var.
func deleteConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	// Get radioID from vars
	vars := mux.Vars(req)
	fmt.Println("ID:", vars["id"])

	// Begin Response
	enableCORS(&w)
	w.WriteHeader(200)
}
