package io

import (
	"io/ioutil"
	"os"

	"github.com/nylo-andry/gowiki"
)

const defaultPath = "pages"

var _ gowiki.PageService = &PageService{}

type PageService struct {
	gowiki.PageService
	Path string
}

func NewPageService() *PageService {
	ps := PageService{}
	ps.Path = defaultPath
	return &ps
}

func (ps *PageService) Save(p *gowiki.Page) error {
	ensurePathExists(ps.Path, 0755)
	filename := p.Title + ".md"
	return ioutil.WriteFile(ps.Path+"/"+filename, []byte(p.Body), 0600)
}

func (ps *PageService) LoadPage(title string) (*gowiki.Page, error) {
	filename := title + ".md"
	body, err := ioutil.ReadFile(ps.Path + "/" + filename)
	if err != nil {
		return nil, err
	}

	return &gowiki.Page{Title: title, Body: string(body)}, nil
}

func ensurePathExists(path string, mode os.FileMode) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, mode)
	}
}
