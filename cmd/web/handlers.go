package main

import (
	"net/http"

	"github.com/gabriellend/try-tarot/pkg/models/cards"
)

type templateParams struct {
	Title string
	Cards []*cards.Card
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	homeParams := templateParams{Title: "Home"}
	err := tpl.ExecuteTemplate(w, "home.page.gohtml", homeParams)
	if err != nil {
		app.serverError(w, err)
	}
}

// learn track
func (app *application) learn(w http.ResponseWriter, r *http.Request) {
	learnParams := templateParams{Title: "Learn"}
	err := tpl.ExecuteTemplate(w, "learn.gohtml", learnParams)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) browse(w http.ResponseWriter, r *http.Request) {
	// ONE WAY TO SHOW THE CARDS
	// store the cards in the repo. you will have to
	// recursively iterate over all the files in ui/static/img/cards
	// and store each file name as a string in a slice, then pass that
	// slice to ExecuteTemplate

	// ANOTHER WAY
	// store them in the database, unmarshal them into card structs,
	// put all of the image fields in a slice and pass that to
	// ExecuteTemplate
	// cards, err := cards.GetAll()

	// browseParams := templateParams{
	// 	Title: "Browse",
	// 	Cards: cards,
	// }

	// // pass that to the template
	// err = tpl.ExecuteTemplate(w, "browse.gohtml", browseParams)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

// read track
func (app *application) read(w http.ResponseWriter, r *http.Request) {
	readParams := templateParams{Title: "Get a reading"}
	err := tpl.ExecuteTemplate(w, "read.gohtml", readParams)
	if err != nil {
		app.serverError(w, err)
	}
}
