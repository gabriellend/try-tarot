package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriellend/try-tarot/pkg/models/cards"
)

type templateData struct {
	View        string
	Cards       []*cards.Card
	CurrentYear int
}

const fileType = ".gohtml"

var functions = template.FuncMap{
	"toBase": toBase,
}

// toBase replaces spaces with hyphens in a string to use it as
// the last element in a url path
func toBase(name string) string {
	b := strings.NewReplacer(" ", "-")
	return strings.ToLower(b.Replace(name))
}

// fromBase reverses toBase. Not currently using.
func fromBase(base string) string {
	n := strings.NewReplacer("-", " ")
	return strings.Title(n.Replace(base))
}

// loadTemplates recursively parses templates and returns them
func loadTemplates(dir string) *template.Template {
	tpl := template.New("").Funcs(functions)

	// Update Walk to WalkDir at some point
	if err := filepath.Walk(
		dir,
		func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, fileType) {
				_, err = tpl.ParseFiles(path)
				if err != nil {
					return err
				}
			}

			return err
		},
	); err != nil {
		log.Fatalln(err)
	}

	return tpl
}
