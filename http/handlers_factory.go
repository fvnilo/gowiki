package http

import (
	"net/http"
)

type handler func(http.ResponseWriter, *http.Request, string)

func makeHandler(fn handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
