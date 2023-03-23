package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/gabriellend/try-tarot/pkg/models/cards"
)

// Initialization
// loadCards takes card information in json format and puts it into a slice of (pointers to) structs
func loadCards() []*cards.Card {
	var cards []*cards.Card

	f, err := os.ReadFile("pkg/models/cards/cards.json")
	if err != nil {
		log.Fatalln(err)
	}

	if err = json.Unmarshal(f, &cards); err != nil {
		log.Fatalln(err)
	}

	return cards
}

// render is a convenience wrapper for ExecuteTemplate
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, data *templateData) {
	err := app.templates.ExecuteTemplate(w, "home.gohtml", data)
	if err != nil {
		app.serverError(w, err)
	}
}

// Errors
// serverError writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

// clientError sends a specific status code and corresponding description
// to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound is a convenience wrapper around clientError which sends a
// 404 Not Found response to the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// perhaps create a function to check errors
