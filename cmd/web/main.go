package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Start server on custom port address (if provided) default is port 3001.
	// Side note: 0-1023 are restricted and reserved for services with root
	// privileges.
	addr := flag.String("addr", ":3001", "HTTP Network Address")

	// Need to parse the command-line arg to actually read in if a user passed
	// a value, and need to parse before it's used anywhere or it'll default
	// to 3001.
	flag.Parse()

	// Create custom loggers to provide standard info.
	// The pipe (|) is used as a join (bitwise OR).
	// NOTE:
	// Custom loggers are also concurrency-safe so it's safe to share them
	// across multiple go-routines or handlers without fear of race conditions.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	// Create a file server for the static files and also make sure
	// it's relative to the current dir path
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Need to strip the "/static" before looking for the files
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// By default http uses the standard logger, so I need to manually use mine
	// Also set the port address and handler to use for the server, allowing me
	// to remove the *addr and mux parameters in the ListenAndServe() call.
	server := http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on http://localhost%s", *addr)
	err := server.ListenAndServe()

	if err != nil {
		errLog.Fatal(err)
	}
}
