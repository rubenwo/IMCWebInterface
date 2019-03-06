package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Credentials is a struct containing a username and password.
// Struct is used to unmarshall.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/login", authenticationEndpoint)
	fmt.Println("Starting authentications server...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

// receives login information in the body of a POST request
// then validates the login credentials and sends a response
func authenticationEndpoint(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In authentication endpoint...")
	var c Credentials
	if err := json.NewDecoder(req.Body).Decode(&c); err != nil {
		w.WriteHeader(422)
		log.Println("Error decoding json to Credentials struct")
		return
	}
	fmt.Println(c)
	uuid, err := authenticate(&c)

	if err != nil {
		log.Println("error authenticating:", err)
	}
	fmt.Println(uuid)

	w.WriteHeader(200)
}

// authenticate a user using the provided credentials
func authenticate(c *Credentials) (string, error) {
	return "", nil
}
