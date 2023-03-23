package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gabriellend/try-tarot/pkg/models/cards"
)

type application struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	cards     []*cards.Card
	templates *template.Template
}

// global variable?
var root = "./ui/templates"

func main() {
	// flags
	port := flag.String("p", ":8080", "port number")
	flag.Parse()

	// loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog:  errorLog,
		infoLog:   infoLog,
		cards:     loadCards(), // "database"
		templates: loadTemplates(root),
	}

	// server
	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *port)
	err := srv.ListenAndServe()
	errorLog.Fatalln(err)
}
