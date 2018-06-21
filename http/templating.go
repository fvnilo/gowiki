package http

import (
	"html/template"
	"net/http"

	"github.com/nylo-andry/gowiki"
)

var templates = template.Must(template.ParseFiles("http/templates/edit.html", "http/templates/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *gowiki.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
