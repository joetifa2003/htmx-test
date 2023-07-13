package renderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
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
	m := minify.New()
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	mw := m.Writer("text/html", w)

	defer func() {
		if err := mw.Close(); err != nil {
			c.Logger().Error(err)
		}
	}()

	switch r := data.(type) {
	case RenderPage:
		var viewContent strings.Builder
		err := t.ExecuteTemplate(&viewContent, name, r.Data)
		if err != nil {
			return err
		}

		return t.ExecuteTemplate(mw, r.Layout, layoutData{
			Title:   r.Title,
			Content: template.HTML(viewContent.String()),
		})
	case RenderWidget:
		return t.ExecuteTemplate(mw, name, r.Data)
	}

	panic("bro why are you here")
}

func NewRenderer() *Renderer {
	return &Renderer{
		template.Must(template.ParseFS(templates, "assets/**/*.html")),
	}
}
