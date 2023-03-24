package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/gabriellend/try-tarot/cmd/web/templates"
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

// Data
// addDefaultData adds common dynamic data to the Data struct passed to templates
func (app *application) addDefaultData(data *templates.Data, r *http.Request) *templates.Data {
	if data == nil {
		data = &templates.Data{}
	}
	data.CurrentYear = time.Now().Year()
	return data
}

// render is a convenience wrapper for ExecuteTemplate
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, data *templates.Data) {
	buf := new(bytes.Buffer)

	err := app.templates.ExecuteTemplate(buf, "home.gohtml", app.addDefaultData(data, r))
	if err != nil {
		app.serverError(w, err)
	}

	buf.WriteTo(w)
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
