package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

var tpl *template.Template

// Make this use the custom loggers?
func init() {
	var err error
	tpl, err = Parse("./ui/templates")
	if err != nil {
		log.Fatalln("Parse error:", err)
	}
}

func main() {
	// flags
	port := flag.String("p", ":8080", "port number")
	flag.Parse()

	// loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
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
