package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	fmt.Println("Starting front-end server on port 80...")
	log.Fatal(http.ListenAndServe(":80", nil))
}
