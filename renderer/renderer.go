package renderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:embed assets/*
var templates embed.FS

type Renderer struct {
	*template.Template
}

type RenderPage struct {
	Title  string
	Layout string
	Data   interface{}
}

type RenderWidget struct {
	Data interface{}
}

type layoutData struct {
	Title   string
	Content template.HTML
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	switch r := data.(type) {
	case RenderPage:
		var viewContent strings.Builder
		err := t.ExecuteTemplate(&viewContent, name, r.Data)
		if err != nil {
			return err
		}

		return t.ExecuteTemplate(w, r.Layout, layoutData{
			Title:   r.Title,
			Content: template.HTML(viewContent.String()),
		})
	case RenderWidget:
		return t.ExecuteTemplate(w, name, r.Data)
	}

	panic("bro why are you here")
}

func NewRenderer() *Renderer {
	return &Renderer{
		template.Must(template.ParseFS(templates, "assets/**/*.html")),
	}
}
