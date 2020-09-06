package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	log.Println("listening to port *:8080. press ctrl + c to cancel.")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("GREETING")
	msg := os.Getenv("MESSAGE")
	anotherMsg := os.Getenv("ANOTHER_MESSAGE")
	fmt.Fprintf(w, `{"env": "%s", "msg", "%s", "anotherMsg": "%s"}`, env, msg, anotherMsg)
}
