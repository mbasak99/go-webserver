package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{"./ui/html/base.html", "./ui/html/pages/home.html", "./ui/html/partials/nav.html"}

	// Read the HTML template into parser
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.errLog.Print(err.Error())
		app.serverError(w, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.errLog.Print(err.Error())
		app.serverError(w, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// Header().Set() method to allow POST calls only
		// Adds "Allow: POST" to the response header map
		// Call this before WriteHeader and/or Write, otherwise it has no effect
		w.Header().Set("Allow", http.MethodPost)
		// Reject any request that's not POST and send a 405 Status (Method not allowed)
		/* w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed")) */
		// http.Error() calls WriteHeader and Write behind the scenes
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	// If you don't call WriteHeader() before calling Write() it automatically
	// sends back a 200 (OK)
	w.Write([]byte("Create a new snippet..."))
}
