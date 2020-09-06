package main

import (
	"fmt"
	"io/ioutil"
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
	content, err := ioutil.ReadFile("/etc/foo/password")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("secret is %q\n", string(content))
	log.Printf("username is %q\n", os.Getenv("USERNAME"))

	fmt.Fprint(w, "hello world")
}
