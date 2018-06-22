package blackfriday

import (
	"html/template"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Renderer struct{}

func (r *Renderer) ToHtml(content string) template.HTML {
	htmlContent := blackfriday.Run([]byte(content))

	return template.HTML(htmlContent)
}
