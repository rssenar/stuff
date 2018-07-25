package main

import (
	"html/template"
	"os"
	"strings"
)

func main() {
	funcMap := template.FuncMap{
		"title":     strings.Title,
		"tableflip": func() string { return "(╯°□°）╯︵ ┻━┻" },
	}

	tpl := template.Must(template.New("main").Funcs(funcMap).ParseGlob("*.html"))
	tplVars := map[string]string{
		"Title":   "hello world",
		"Content": "Hi there",
	}
	tpl.ExecuteTemplate(os.Stdout, "index.html", tplVars)
}
