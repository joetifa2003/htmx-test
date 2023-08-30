package renderer

import (
	"embed"

	"github.com/a-h/templ"
	"github.com/joetifa2003/htmx-test/renderer/templates/widgets"
	"github.com/labstack/echo/v4"
)

//go:embed assets/*
var Assets embed.FS

type customWidget struct {
	id     string
	widget templ.Component
}

type Render struct {
	Component     templ.Component
	customWidgets []customWidget
}

func (t *Render) AddCustomWidget(id string, widget templ.Component) {
	t.customWidgets = append(t.customWidgets, customWidget{
		id:     id,
		widget: widget,
	})
}

func (t *Render) Render(c echo.Context) error {
	ctx := c.Request().Context()
	w := c.Response().Writer
	for _, v := range t.customWidgets {
		err := widgets.HtmxSwap(v.id, v.widget).Render(ctx, w)

		if err != nil {
			return err
		}
	}

	return t.Component.Render(ctx, w)
}
