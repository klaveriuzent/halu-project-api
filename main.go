package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//sample http server for testing purposes
	http.HandleFunc("/", ExampleHandler)
	log.Println("** Service Started on Port 8080 **")
	err := http.ListenAndServe("0.0.0.0::8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
