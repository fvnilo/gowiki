package http

import (
	"net/http"

	"github.com/nylo-andry/gowiki"
)

func save(w http.ResponseWriter, r *http.Request, title string, ps gowiki.PageService) {
	body := r.FormValue("body")
	p := &gowiki.Page{Title: title, Body: []byte(body)}
	err := ps.Save(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func view(w http.ResponseWriter, r *http.Request, title string, ps gowiki.PageService) {
	p, err := ps.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func edit(w http.ResponseWriter, r *http.Request, title string, ps gowiki.PageService) {
	p, err := ps.LoadPage(title)
	if err != nil {
		p = &gowiki.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
