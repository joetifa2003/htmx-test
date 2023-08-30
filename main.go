package main

import (
	"strconv"

	"github.com/joetifa2003/htmx-test/renderer"
	"github.com/joetifa2003/htmx-test/renderer/templates/layouts"
	"github.com/joetifa2003/htmx-test/renderer/templates/pages"
	"github.com/joetifa2003/htmx-test/renderer/templates/widgets"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.StaticFS("/", renderer.Assets)

	// for the main page
	e.GET("/", func(c echo.Context) error {
		rndr := renderer.Render{}
		rndr.Component = layouts.MainLayout("index page", pages.IndexPage())
		return rndr.Render(c)
	})

	// for the counter widget /renderer/assets/widgets/counter.html
	e.POST("/counter", func(c echo.Context) error {
		rndr := renderer.Render{}
		count := c.FormValue("count")
		countInt, err := strconv.Atoi(count)
		if err != nil {
			return err
		}
		countInt++
		rndr.Component = widgets.Counter(countInt)
		rndr.AddCustomWidget("counter-display-wrapper", widgets.CounterDisplay(countInt))
		return rndr.Render(c)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
