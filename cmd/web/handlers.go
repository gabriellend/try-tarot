package main

import (
	"fmt"
	"io"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

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
	card := r.URL.Query().Get("card")

	io.WriteString(w, fmt.Sprintf("the card is %v", card))
	// data := templateData{
	// 	View:  ,
	// 	Cards: cards,
	// }

	// err = tpl.ExecuteTemplate(w, "browse.gohtml", data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
}

// read track
func (app *application) read(w http.ResponseWriter, r *http.Request) {
	data := &templateData{View: "Get a reading"}

	app.render(w, r, "read.gohtml", data)
}
