package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	UID   string `json:"uid"`
	Error string `json:"error"`
}

type Accounts struct {
	Credentials []Credentials `json:"accounts"`
}

// Credentials is a struct containing a username and password.
// Struct is used to unmarshall.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var accounts Accounts

func main() {
	initAccounts()
	http.HandleFunc("/login", authenticationEndpoint)
	fmt.Println("Starting authentications server...")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func initAccounts() {
	fmt.Println("Loading accounts...")
	jsonFile, err := os.Open("accounts.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened accounts.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &accounts)
	fmt.Println("Unmarshalled accounts.json to accounts variable")
	for i := 0; i < len(accounts.Credentials); i++ {
		fmt.Printf("[LOADED] Username: %s, Password: %s\n", accounts.Credentials[i].Username, accounts.Credentials[i].Password)
	}
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
	fmt.Println("UUID:", uuid)

	response := Response{
		UID:   uuid,
		Error: fmt.Sprint(err)}

	resp, err := json.Marshal(response)
	if err != nil {
		log.Println("error marshalling:", err)
	}
	w.Write(resp)
}

// authenticate a user using the provided credentials
func authenticate(c *Credentials) (string, error) {
	for i := 0; i < len(accounts.Credentials); i++ {
		username := accounts.Credentials[i].Username
		password := accounts.Credentials[i].Password
		if username == c.Username && password == c.Password {
			return "Authenticated", nil
		}
	}
	return "", errors.New("Couldn't authenticate the user")
}
