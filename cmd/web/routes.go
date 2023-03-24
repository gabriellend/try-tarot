package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	// learn track
	mux.Get("/learn", http.HandlerFunc(app.learn))
	mux.Get("/learn/browse", http.HandlerFunc(app.browse))
	mux.Get("/learn/browse/:card", http.HandlerFunc(app.showCard))
	// read track
	mux.Get("/read", http.HandlerFunc(app.read))
	// images, css, js
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
