package main

import (
	"net/http"

	"github.com/gabriellend/try-tarot/pkg/models/cards"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := &templateData{View: "Home"}

	app.render(w, r, "home.gohtml", data)
}

// learn track
func (app *application) learn(w http.ResponseWriter, r *http.Request) {
	data := &templateData{View: "Learn"}

	app.render(w, r, "learn.gohtml", data)
}

func (app *application) browse(w http.ResponseWriter, r *http.Request) {
	data := &templateData{
		View:  "Browse",
		Cards: app.cards,
	}

	app.render(w, r, "browse.gohtml", data)
}

func (app *application) showCard(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(":card")
	var card []*cards.Card

	for _, v := range app.cards {
		if toBase(v.Name) == name {
			card = append(card, v)
		}
	}

	data := &templateData{
		View:  "Browse",
		Cards: card,
	}

	app.render(w, r, "card.gohtml", data)
}

// read track
func (app *application) read(w http.ResponseWriter, r *http.Request) {
	data := &templateData{View: "Get a reading"}

	app.render(w, r, "read.gohtml", data)
}
