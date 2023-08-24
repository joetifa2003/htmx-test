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
var Assets embed.FS

type Renderer struct {
	*template.Template

	defaultLayout string
}

type RenderWidget struct {
	Data interface{}
}

type RenderPage struct {
	RenderWidget

	Title  string
	Layout string
}

type layoutData struct {
	Title   string
	Content template.HTML
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepDefaultAttrVals: true,
	})
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

		layout := r.Layout
		if layout == "" {
			layout = t.defaultLayout
		}

		return t.ExecuteTemplate(mw, layout, layoutData{
			Title:   r.Title,
			Content: template.HTML(viewContent.String()),
		})
	case RenderWidget:
		return t.ExecuteTemplate(mw, name, r.Data)
	}

	panic("should pass RenderPage or RenderWidget structs to the data parameter")
}

func NewRenderer(defaultLayout string) *Renderer {
	return &Renderer{
		Template:      template.Must(template.ParseFS(Assets, "assets/**/*.html")),
		defaultLayout: defaultLayout,
	}
}
