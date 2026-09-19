package main

import (
	"fmt"
	"log"
	"net/http"
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func handler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	// w.Header().Set("Content-Type", "application/json")
	// w.Write({test: "This is a test json object"})
	fmt.Fprintf(w, "Hello there! This is the %s route.", r.URL.Path[1:])
}

func main() {
	fmt.Println("Server listening on port 8080.")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
