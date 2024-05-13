package main

import (
	"flag"
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

	// Start the server on custom port address (if provided) default is 3001
	addr := flag.String("addr", ":3001", "HTTP Network Address")

	log.Printf("Start server on http://localhost%s", *addr)
	err := http.ListenAndServe(*addr, mux)

	if err != nil {
		log.Fatal(err)
	}
}
