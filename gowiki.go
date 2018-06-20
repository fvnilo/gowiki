package gowiki

// Page represents the fields composing a wiki page.
type Page struct {
	Title string
	Body  []byte
}

// PageService is the interface for manipulating Pages
type PageService interface {
	LoadPage(title string) (*Page, error)
	Save(p *Page) error
}
