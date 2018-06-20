package http

import (
	"net/http"

	"github.com/nylo-andry/gowiki"
)

type handler func(http.ResponseWriter, *http.Request, string, gowiki.PageService)

func makeHandler(fn handler, ps gowiki.PageService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2], ps)
	}
}
