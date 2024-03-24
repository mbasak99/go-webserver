package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Create a file server for the static files and also make sure
	// it's relative to the current dir path
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Need to strip the "/static" before looking for the files
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Start server on localhost:3001")
	err := http.ListenAndServe(":3001", mux)

	if err != nil {
		log.Fatal(err)
	}
}
