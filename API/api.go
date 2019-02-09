package API

import (
	"IMCWebInterface/IMC"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Port    string
	Running bool
}

type response struct {
	Data string `json:"data,omitempty"`
}

type Credentials struct {
	Username string `json:"username, omitempty"`
	Password string `json:"password, omitempty"`
}

type loginResponse struct {
	Valid bool   `json:"valid,omitempty"`
	UID   string `json:"uid,omitempty"`
}

var imc *IMC.TcpIMCServer

func (s *Server) Start(i *IMC.TcpIMCServer) error {
	if s.Port == "" {
		s.Port = "8081"
	}
	s.Running = true
	if i.Running == false {
		return errors.New("IMC server is not running...\n")
	}
	imc = i
	router := mux.NewRouter()
	router.HandleFunc("/api/config", updateConfigEndpoint).Methods("POST")
	router.HandleFunc("/api/config", getConfigEndpoint).Methods("GET")
	router.HandleFunc("/api/radio", getRadioEndpoint).Methods("GET")
	router.HandleFunc("/login", loginEndpoint).Methods("POST")
	http.ListenAndServe(":"+s.Port, router)
	return nil
}
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func updateConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	enableCORS(&w)
	d, err := json.Marshal(response{"Update Config Endpoint"})
	if err != nil {
		panic(err)
	}
	w.Write(d)
}

func getConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	enableCORS(&w)
	d, err := json.Marshal(response{"Get Config Endpoint"})
	if err != nil {
		panic(err)
	}
	w.Write(d)
}

func getRadioEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)
	enableCORS(&w)
	d, err := json.Marshal(response{"Get Radio Endpoint"})
	if err != nil {
		panic(err)
	}
	w.Write(d)
}
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

func validate(credentials *Credentials) loginResponse {
	return loginResponse{
		Valid: true,
		UID:   "12345-abcde-54321"}
}
