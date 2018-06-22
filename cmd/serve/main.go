package main

import (
	"github.com/nylo-andry/gowiki/blackfriday"
	"github.com/nylo-andry/gowiki/http"
	"github.com/nylo-andry/gowiki/io"
)

func main() {
	ps := io.NewPageService()
	r := &blackfriday.Renderer{}
	s := http.NewServer(ps, r)

	s.Serve()
}
