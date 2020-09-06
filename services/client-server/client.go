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
	svc := os.Getenv("SERVER_SERVICE_URL")
	log.Println(svc)
	resp, err := http.Get(svc)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, string(body))
}
