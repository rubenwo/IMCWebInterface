package FrontEnd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Port string
}

type Credentials struct {
	Username string `json:"username, omitempty"`
	Password string `json:"password, omitempty"`
}

type loginResponse struct {
	Valid bool   `json:"valid,omitempty"`
	UID   string `json:"uid,omitempty"`
}

var fs = http.FileServer(http.Dir("FrontEnd/public"))

func (s *Server) Start() {
	if s.Port == "" {
		s.Port = "80"
	}
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login", loginEndpoint)
	http.ListenAndServe(":"+s.Port, nil)
}
func rootHandler(w http.ResponseWriter, req *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fs.ServeHTTP(w, req)
}

func loginEndpoint(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
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
		break
	case "GET":
		fmt.Println("Got GET request")
		break
	}
}

func validate(credentials *Credentials) loginResponse {
	return loginResponse{
		Valid: true,
		UID:   "12345-abcde-54321"}
}
