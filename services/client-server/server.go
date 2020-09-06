package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	log.Println("listening to port *:8080. press ctrl + c to cancel.")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", mux))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("calling server")
	fmt.Fprint(w, `{"hello": "world"}`)
}
