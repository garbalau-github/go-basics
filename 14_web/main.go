package main

import (
	"fmt" 
	"net/http"
)

// ~ Pages
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Golang!</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	// Router
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server starting...")

	// Listener
	http.ListenAndServe(":3000", nil)
}