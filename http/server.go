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
	s := &Server{defaultPort}

	http.HandleFunc("/view/", makeHandler(view, ps))
	http.HandleFunc("/edit/", makeHandler(edit, ps))
	http.HandleFunc("/save/", makeHandler(save, ps))

	return s
}

// Start will serve the application.
func (ph *Server) Start() {
	log.Printf("About to listen on port: %v", ph.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", ph.Port), nil))
}
