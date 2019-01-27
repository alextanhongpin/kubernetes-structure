package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	var (
		serviceURI = flag.String("service_uri", "http://server", "the service endpoint")
		host       = flag.String("host", "0.0.0.0", "the host")
		port       = flag.String("port", "8080", "the port")
		text       = flag.String("text", "hello", "the text to distinguish which server is it")
	)
	flag.Parse()

	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		resp, err := client.Get(*serviceURI)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(body))
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, *text)
	})
	addr := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Printf("listening to %s\n", addr)
	http.ListenAndServe(addr, nil)

}
