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

func loadTemplates(dir string) *template.Template {
	tpl := template.New("")

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
