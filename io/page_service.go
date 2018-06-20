package io

import (
	"io/ioutil"

	"github.com/nylo-andry/gowiki"
)

type PageService struct{ gowiki.PageService }

func (ps *PageService) Save(p *gowiki.Page) error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func (ps *PageService) LoadPage(title string) (*gowiki.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &gowiki.Page{Title: title, Body: body}, nil
}
