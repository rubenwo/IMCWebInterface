package API

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Port string
}

func (s *Server) Start() {
	if s.Port == "" {
		s.Port = "8081"
	}
	router := mux.NewRouter()
	router.HandleFunc("/api/config", updateConfigEndpoint).Methods("POST")
	router.HandleFunc("/api/config", getConfigEndpoint).Methods("GET")
	router.HandleFunc("/api/radio", getRadioEndpoint).Methods("GET")
	http.ListenAndServe(":"+s.Port, router)
}

func updateConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	w.Write([]byte("You have reached the Update Config Endpoint"))
}

func getConfigEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RemoteAddr)
	w.Write([]byte("You have reached the Get Config Endpoint"))
}

func getRadioEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)
	w.Write([]byte("You have reached Get Radio Endpoint"))
}
