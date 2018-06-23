package main

import (
	"github.com/nylo-andry/gowiki/http"
	"github.com/nylo-andry/gowiki/io"
)

func main() {
	ps := io.NewPageService()
	s := http.NewServer(ps)

	s.Serve()
}
