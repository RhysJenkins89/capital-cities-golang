package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there! This is the %s route.", r.URL.Path[1:])
}

func main() {
	fmt.Println("Server listening on port 8080.")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))	
}

