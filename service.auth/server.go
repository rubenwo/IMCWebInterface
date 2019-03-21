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
	UID     string `json:"uid"`
	Devices []int  `json:"devices"`
	Error   string `json:"error"`
}

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Devices  []int  `json:"devices"`
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
	http.HandleFunc("/auth", authenticationEndpoint)
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
	for i := 0; i < len(accounts.Accounts); i++ {
		fmt.Printf("[LOADED] Username: %s, Password: %s\n", accounts.Accounts[i].Username, accounts.Accounts[i].Password)
	}
}

// Enables CORS for http responses
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
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
	uuid, devices, err := authenticate(&c)

	if err != nil {
		log.Println("error authenticating:", err)
		w.WriteHeader(401)
		return
	}
	fmt.Println("UUID:", uuid)

	response := Response{
		UID:     uuid,
		Devices: devices,
		Error:   ""}

	resp, err := json.Marshal(response)
	if err != nil {
		log.Println("error marshalling:", err)
		w.WriteHeader(500)
		return
	}
	enableCORS(&w)
	w.Write(resp)
}

// authenticate a user using the provided credentials
func authenticate(c *Credentials) (string, []int, error) {
	for i := 0; i < len(accounts.Accounts); i++ {
		username := accounts.Accounts[i].Username
		password := accounts.Accounts[i].Password
		if username == c.Username && password == c.Password {
			return "Authenticated", accounts.Accounts[i].Devices, nil
		}
	}
	return "", nil, errors.New("couldn't authenticate the user")
}
