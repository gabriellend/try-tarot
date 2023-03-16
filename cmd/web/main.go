package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = Parse("./ui/templates")
	if err != nil {
		log.Fatalln("Parse error:", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// learn track
	mux.HandleFunc("/learn", learn)
	mux.HandleFunc("/learn/browse", browse)

	// read track
	mux.HandleFunc("/read", read)

	// images
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatalln(err)
}
