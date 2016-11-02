// render.go
// Copyright (C) 2016 Reza Jatnika <rezajatnika@gmail.com>
//
// Distributed under terms of the MIT license.
//

package views

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

var (
	layoutPath string
	templates  *template.Template
	tmplPath   string
)

func init() {
	tmplPath = path.Dir("app/views/")
	layoutPath = tmplFile("layout")
}

// Render renders template
func Render(w http.ResponseWriter, tmpl string, data interface{}) error {
	tmpl = tmplFile(tmpl)

	// Parse template files
	t, err := templates.ParseFiles(layoutPath, tmpl)
	checkErr(err)

	// Set HTML header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Execute template
	return t.ExecuteTemplate(w, "layout", data)
}

func tmplFile(tmpl string) string {
	return tmplPath + "/" + tmpl + ".tmpl"
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
