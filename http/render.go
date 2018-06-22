package http

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))

func renderTemplate(w http.ResponseWriter, templateName string, content interface{}) {
	err := templates.ExecuteTemplate(w, templateName+".html", content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
