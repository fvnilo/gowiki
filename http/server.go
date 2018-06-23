package http

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/nylo-andry/gowiki"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

const defaultPort = 8080

// Server represents the mecanism that handles Page requests
type Server struct {
	Port int
}

// NewServer creates a new Server
func NewServer(ps gowiki.PageService) *Server {
	pc := newPageController(ps)
	s := &Server{defaultPort}

	http.HandleFunc("/view/", makeHandler(pc.view))
	http.HandleFunc("/edit/", makeHandler(pc.edit))
	http.HandleFunc("/save/", makeHandler(pc.save))

	return s
}

// Start will serve the application.
func (s *Server) Serve() {

	log.Printf("About to listen on port: %v", s.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", s.Port), nil))
}
