package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	err := tpl.ExecuteTemplate(w, "home.page.gohtml", "Home")
	if err != nil {
		app.serverError(w, err)
	}
}

// learn track
func (app *application) learn(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "learn.gohtml", "Learn")
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) browse(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "browse.gohtml", "Browse")
	if err != nil {
		app.serverError(w, err)
	}
}

// read track
func (app *application) read(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "read.gohtml", "Get a reading")
	if err != nil {
		app.serverError(w, err)
	}
}
