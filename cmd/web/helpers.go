package main

import (
	"log"
	"os"
	"path/filepath"
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
