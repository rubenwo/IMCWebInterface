package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	/* Testing the service.auth from another image
	resp, err := http.Post(
		"http://service.auth/login",
		"application/json",
		bytes.NewBuffer([]byte(`"username":"Ruben","password":"test"`)))
	fmt.Println(resp.Status)
	fmt.Println(err)
	*/
	fmt.Println("Starting front-end server on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
