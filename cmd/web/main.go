package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Start server on custom port address (if provided) default is port 3001.
	addr := flag.String("addr", ":3001", "HTTP Network Address")

	// Need to parse the command-line arg to actually read in if a user passed
	// a value, and need to parse before it's used anywhere or it'll default
	// to 3001.
	flag.Parse()

	mux := http.NewServeMux()

	// Create a file server for the static files and also make sure
	// it's relative to the current dir path
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Need to strip the "/static" before looking for the files
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Start server on http://localhost%s", *addr)
	err := http.ListenAndServe(*addr, mux)

	if err != nil {
		log.Fatal(err)
	}
}
