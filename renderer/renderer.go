package renderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

//go:embed assets/*
var templates embed.FS

type Renderer struct {
	*template.Template
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.ExecuteTemplate(w, name, data)
}

func NewRenderer() *Renderer {
	return &Renderer{
		template.Must(template.ParseFS(templates, "assets/**/*.html")),
	}
}
