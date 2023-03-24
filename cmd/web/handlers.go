package main

import (
	"net/http"

	"github.com/gabriellend/try-tarot/cmd/web/templates"
	"github.com/gabriellend/try-tarot/pkg/models/cards"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := &templates.Data{View: "Home"}

	app.render(w, r, "home.gohtml", data)
}

// learn track
func (app *application) learn(w http.ResponseWriter, r *http.Request) {
	data := &templates.Data{View: "Learn"}

	app.render(w, r, "learn.gohtml", data)
}

func (app *application) browse(w http.ResponseWriter, r *http.Request) {
	data := &templates.Data{
		View:  "Browse",
		Cards: app.cards,
	}

	app.render(w, r, "browse.gohtml", data)
}

func (app *application) showCard(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(":card")
	card, err := cards.Get(name, app.cards)
	if err != nil {
		app.serverError(w, err)
	}

	data := &templates.Data{
		View:  "Browse",
		Cards: card,
	}

	app.render(w, r, "card.gohtml", data)
}

// read track
func (app *application) read(w http.ResponseWriter, r *http.Request) {
	data := &templates.Data{View: "Get a reading"}

	app.render(w, r, "read.gohtml", data)
}
