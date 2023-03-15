package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	// Maybe create a helper for this this
	tpl = template.New("")
	if err := filepath.Walk(
		"./ui/templates",
		func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, ".gohtml") {
				_, err = tpl.ParseFiles(path)
				if err != nil {
					log.Fatalln(err)
				}
			}

			return err
		},
	); err != nil {
		log.Fatalln("failed to parse files:", err)
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

	// Remove this when you have a favicon
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatalln(err)
}
