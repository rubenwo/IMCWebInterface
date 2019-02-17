package API

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Start's the REST API server
func (s *Server) Start() {
	if s.Port == "" {
		s.Port = "8081"
	}
	s.Running = true

	router := mux.NewRouter()
	router.HandleFunc("/api/config/{radioID}", updateConfigEndpoint).Methods("POST")
	router.HandleFunc("/api/config/{radioID}", getConfigEndpoint).Methods("GET")
	router.HandleFunc("/api/radios/{accountID}", getRadiosEndpoint).Methods("GET")
	router.HandleFunc("/login", loginEndpoint).Methods("POST")
	fmt.Println("API Server is running...")
	log.Fatal(http.ListenAndServe(":"+s.Port, router))
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
	d, err := json.Marshal(response{"Update Config Endpoint"})
	if err != nil {
		panic(err)
	}
	w.Write(d)
}

// GET: Get the config of 1 radio. Has a radioID as var.
func getConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	// Get radioID from vars
	vars := mux.Vars(req)
	fmt.Println("Radio ID:", vars["radioID"])

	// Begin Response
	enableCORS(&w)
	d, err := json.Marshal(response{"Get Config Endpoint"})
	if err != nil {
		panic(err)
	}
	w.Write(d)
}

// GET: Get all radio saved to an account UID. Has accountUID as var.
func getRadiosEndpoint(w http.ResponseWriter, req *http.Request) {
	// Get radioID from vars
	vars := mux.Vars(req)
	fmt.Println("Account ID:", vars["accountID"])

	// Begin Response
	enableCORS(&w)
	d, err := json.Marshal(response{"Get Radio Endpoint"})
	if err != nil {
		panic(err)
	}
	w.Write(d)
}

// POST: Login endpoint. Receives JSON Credentials object as body.
func loginEndpoint(w http.ResponseWriter, req *http.Request) {
	enableCORS(&w)

	var c Credentials
	if err := json.NewDecoder(req.Body).Decode(&c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
	resp, err := json.Marshal(validate(&c))
	if err != nil {
		panic(err)
	}
	w.Write(resp)
}

// Validate login based on credentials. Returns a login response.
func validate(credentials *Credentials) loginResponse {
	return loginResponse{
		Valid: true,
		UID:   "12345-abcde-54321"}
}
