package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)
//Configs is a struct used for decoding the configs.json file
type Configs struct {
	Configs []Config `json:"configs"`
}

var configs Configs

func main() {
	loadConfigs()
	router := mux.NewRouter()

	router.HandleFunc("/config", updateConfigEndpoint).Methods("POST")
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

// POST: Updates the config of 1 radio. Receives JSON Config object as body.
func updateConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	// Decode Config from body
	var newConfig Config
	if err := json.NewDecoder(req.Body).Decode(&newConfig); err != nil {
		log.Fatal(err)
	}
	fmt.Println("New Config:", newConfig)
	config, err := lookupConfig(newConfig.ID)
	if err != nil {
		log.Println(err)
		configs.Configs = append(configs.Configs, newConfig)
		fmt.Println("Added New config to configs")
	} else {
		*config = newConfig
		fmt.Println("Updated config with ID:", config.ID)
	}
	saveConfigs()
	// Begin Response
	enableCORS(&w)
	w.WriteHeader(200)
}

// GET: Get the config of 1 radio. Has a radioID as var.
func getConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	// Get radioID from vars
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Println("ID:", id)
	config, err := lookupConfig(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}
	fmt.Println(config)
	// Begin Response

	data, err := json.Marshal(config)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}
	enableCORS(&w)
	w.Write(data)
}

// GET: Get the config of 1 radio. Has a radioID as var.
func deleteConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	// Get radioID from vars
	vars := mux.Vars(req)
	id := vars["id"]
	fmt.Println("ID:", id)
	config, err := lookupConfig(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(404)
		return
	}
	fmt.Println(config)
	for i := 0; i < len(configs.Configs); i++ {
		if config == &configs.Configs[i] {
			configs.Configs = append(configs.Configs[:i], configs.Configs[i+1:]...)
		}
	}

	// Begin Response
	enableCORS(&w)
	w.WriteHeader(200)
}

func lookupConfig(id string) (*Config, error) {
	for i := 0; i < len(configs.Configs); i++ {
		if configs.Configs[i].ID == id {
			return &configs.Configs[i], nil
		}
	}
	return nil, errors.New("no config found with that ID")
}

func saveConfigs() {

}
