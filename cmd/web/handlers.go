package main

import (
	"log"
	"net/http"
)

type pageInfo struct {
	Title   string
	Header  string
	ShowNav bool
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	homeInfo := pageInfo{
		Title:   "Home",
		Header:  "Welcome",
		ShowNav: false,
	}

	err := tpl.ExecuteTemplate(w, "home.page.gohtml", homeInfo)
	if err != nil {
		log.Println("template execution failed in home:", err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// learn track
func learn(w http.ResponseWriter, r *http.Request) {
	learnInfo := pageInfo{
		Title:   "Learn",
		Header:  "Learn",
		ShowNav: true,
	}
	err := tpl.ExecuteTemplate(w, "learn.gohtml", learnInfo)
	if err != nil {
		log.Println("template execution failed in learn:", err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func browse(w http.ResponseWriter, r *http.Request) {
	browseInfo := pageInfo{
		Title:   "Browse",
		Header:  "Browse",
		ShowNav: true,
	}

	err := tpl.ExecuteTemplate(w, "browse.gohtml", browseInfo)
	if err != nil {
		log.Println("template execution failed in browse:", err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// read track
func read(w http.ResponseWriter, r *http.Request) {
	readInfo := pageInfo{
		Title:   "Get a reading",
		Header:  "Ger a reading",
		ShowNav: true,
	}

	err := tpl.ExecuteTemplate(w, "read.gohtml", readInfo)
	if err != nil {
		log.Println("template execution failed in read:", err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
