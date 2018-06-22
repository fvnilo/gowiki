package http

import (
	"html/template"
	"net/http"

	"github.com/nylo-andry/gowiki"
)

type pageController struct {
	pageService gowiki.PageService
	renderer    gowiki.Renderer
}

type ViewModel struct {
	Title string
	Body  template.HTML
}

func newPageController(ps gowiki.PageService, r gowiki.Renderer) *pageController {
	return &pageController{ps, r}
}

func (h *pageController) save(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &gowiki.Page{Title: title, Body: body}
	err := h.pageService.Save(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func (h *pageController) view(w http.ResponseWriter, r *http.Request, title string) {
	p, err := h.pageService.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	body := h.renderer.ToHtml(p.Body)
	renderTemplate(w, "view", ViewModel{p.Title, body})
}

func (h *pageController) edit(w http.ResponseWriter, r *http.Request, title string) {
	p, err := h.pageService.LoadPage(title)
	if err != nil {
		p = &gowiki.Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}
