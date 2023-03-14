package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	// tpl = template.Must(template.ParseGlob("templates/*"))
	// tpl = template.New("")
	// if err := filepath.WalkDir(
	// 	"./templates",
	// 	func(path string, info os.FileInfo, err error) error {
	// 		if strings.Contains(path, ".gohtml") {
	// 			_, err = tpl.ParseFiles(path)
	// 			if err != nil {
	// 				log.Fatalln(err)
	// 			}
	// 		}

	// 		return err
	// 	},
	// ); err != nil {
	// 	log.Fatalln("failed to parse files")
	// }
}

func main() {
	http.HandleFunc("/", index)

	// learn
	http.HandleFunc("/learn", learn)
	http.HandleFunc("/learn/browse", browse)

	// read
	http.HandleFunc("/read", read)

	// Remove this when you have a favicon
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func learn(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "learn/learn.gohtml", nil)
}

func browse(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "learn/browse.gohtml", nil)
}

func read(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "read.gohtml", nil)
}
