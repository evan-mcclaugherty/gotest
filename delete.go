package main

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "<h1>Welcome to the Go toolset.</h1>")

}

func main() {

	fmt.Println("Hello.")
	fmt.Println("Starting http server.")
	// Register handler function
	http.HandleFunc("/welcome", Welcome)
	fmt.Println("Go to localhost:8080/welcome")
	fmt.Println("To terminate press CTRL+C.")
	// Start server
	http.ListenAndServe(":8080", nil)

}
