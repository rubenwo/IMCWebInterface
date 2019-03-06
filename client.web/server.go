package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	/* Testing the service.config from another image
	resp, err := http.Get("http://service.config/config/1")
	fmt.Println(resp.Status)
	fmt.Println(err)
	*/
	fmt.Println("Starting front-end server on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
