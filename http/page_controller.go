package http

import (
	"html/template"
	"net/http"

	"github.com/nylo-andry/gowiki"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type pageController struct {
	pageService gowiki.PageService
}

type ViewModel struct {
	Title string
	Body  template.HTML
}

func newPageController(ps gowiki.PageService) *pageController {
	return &pageController{ps}
}

func (pc *pageController) save(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &gowiki.Page{Title: title, Body: body}
	err := pc.pageService.Save(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func (pc *pageController) view(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pc.pageService.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	body := toHTML(p.Body)
	renderTemplate(w, "view", ViewModel{p.Title, body})
}

func (pc *pageController) edit(w http.ResponseWriter, r *http.Request, title string) {
	p, err := pc.pageService.LoadPage(title)
	if err != nil {
		p = &gowiki.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func toHTML(markdown string) template.HTML {
	htmlContent := blackfriday.Run([]byte(markdown))
	return template.HTML(htmlContent)
}
