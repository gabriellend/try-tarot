package templates

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriellend/try-tarot/pkg/models/cards"
)

type Data struct {
	View        string
	Cards       []*cards.Card
	CurrentYear int
	Set         string
}

const fileType = ".gohtml"

var funcs = template.FuncMap{
	"ToBase": cards.ToBase,
	"Random": cards.Random,
}

// fromBase reverses ToBase. Not currently using.
// func FromBase(base string) string {
// 	n := strings.NewReplacer("-", " ")
// 	return strings.Title(n.Replace(base))
// }

// loadTemplates recursively parses templates and returns them
func LoadTemplates(dir string) *template.Template {
	tpl := template.New("").Funcs(funcs)

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
