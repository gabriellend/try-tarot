package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := tpl.ExecuteTemplate(w, "home.page.gohtml", "Home")
	if err != nil {
		log.Println("template execution failed in home:", err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// learn track
func learn(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "learn.gohtml", "Learn")
	if err != nil {
		log.Println("template execution failed in learn:", err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func browse(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "browse.gohtml", "Browse")
	if err != nil {
		log.Println("template execution failed in browse:", err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// read track
func read(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "read.gohtml", "Get a reading")
	if err != nil {
		log.Println("template execution failed in read:", err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
