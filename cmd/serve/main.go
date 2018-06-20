package main

import (
	"github.com/nylo-andry/gowiki/http"
	"github.com/nylo-andry/gowiki/io"
)

func main() {
	ps := &io.PageService{}
	s := http.NewServer(ps)

	s.Start()
}
