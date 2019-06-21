package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load the environment variable.
	greeting := os.Getenv("GREETING")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"message": "%s"}`, greeting)
	})
	log.Println("listening to port *:8080)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}
