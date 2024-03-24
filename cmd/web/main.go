package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Start server on localhost:3001")
	err := http.ListenAndServe(":3001", mux)

	if err != nil {
		log.Fatal(err)
	}
}
