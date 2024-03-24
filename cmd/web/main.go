package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Header().Set() method to allow POST calls only
		// Adds "Allow: POST" to the response header map
		// Call this before WriteHeader and/or Write, otherwise it has no effect
		w.Header().Set("Allow", http.MethodPost)
		// Reject any request that's not POST and send a 405 Status (Method not allowed)
		/* w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed")) */
		// http.Error() calls WriteHeader and Write behind the scenes
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	// If you don't call WriteHeader() before calling Write() it automatically
	// sends back a 200 (OK)
	w.Write([]byte("Create a new snippet..."))
}

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
