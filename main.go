package main

import (
	"log"
	"net/http"

	"github.com/nylo-andry/gowiki/wiki"
)

func main() {
	http.HandleFunc("/view/", wiki.ViewHandler)
	http.HandleFunc("/edit/", wiki.EditHandler)
	http.HandleFunc("/save/", wiki.SaveHandler)

	log.Println("About to listen on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
