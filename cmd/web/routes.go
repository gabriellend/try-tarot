package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	// learn track
	mux.HandleFunc("/learn", app.learn)
	mux.HandleFunc("/learn/browse", app.browse)
	mux.HandleFunc("/learn/browse/card", app.showCard)

	// read track
	mux.HandleFunc("/read", app.read)

	// images
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
