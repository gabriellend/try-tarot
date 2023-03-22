package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gabriellend/try-tarot/pkg/models/cards"
)

type templateParams struct {
	View  string
	Cards []cards.Card
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	homeParams := templateParams{View: "Home"}
	err := tpl.ExecuteTemplate(w, "home.gohtml", homeParams)
	if err != nil {
		app.serverError(w, err)
	}
}

// learn track
func (app *application) learn(w http.ResponseWriter, r *http.Request) {
	learnParams := templateParams{View: "Learn"}
	err := tpl.ExecuteTemplate(w, "learn.gohtml", learnParams)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) browse(w http.ResponseWriter, r *http.Request) {
	var cards []cards.Card

	f, err := os.ReadFile("pkg/models/cards/cards.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(f, &cards)

	browseParams := templateParams{
		View:  "Browse",
		Cards: cards,
	}

	// // pass that to the template
	err = tpl.ExecuteTemplate(w, "browse.gohtml", browseParams)
	if err != nil {
		app.serverError(w, err)
	}
}

// read track
func (app *application) read(w http.ResponseWriter, r *http.Request) {
	readParams := templateParams{View: "Get a reading"}
	err := tpl.ExecuteTemplate(w, "read.gohtml", readParams)
	if err != nil {
		app.serverError(w, err)
	}
}
