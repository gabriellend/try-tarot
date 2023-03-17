package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"text/template"
)

// Templates
const fileType = ".gohtml"

func Parse(dir string) (*template.Template, error) {
	tpl := template.New("")
	// Update Walk to WalkDir at some point
	if err := filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, fileType) {
				_, err = tpl.ParseFiles(path)
				if err != nil {
					log.Println("ParseFiles error:", err)
				}
			}

			return err
		},
	); err != nil {
		return nil, err
	}

	return tpl, nil
}

// Errors
// serverError writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

// clientError sends a specific status code and corresponding description
// to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound is a convenience wrapper around clientError which sends a
// 404 Not Found response to the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
