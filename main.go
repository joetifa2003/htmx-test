package main

import (
	"net/http"
	"strconv"

	"github.com/joetifa2003/htmx-test/renderer"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = renderer.NewRenderer() // sets the renderer, which executes templates

	// on dev mode, templates are reparsed on every request (hot reloading)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			e.Renderer = renderer.NewRenderer()
			next(c)
			return nil
		}
	})

	// for the counter widget /renderer/assets/widgets/counter.html
	e.POST("/counter", func(c echo.Context) error {
		count := c.FormValue("count")
		countInt, err := strconv.Atoi(count)
		if err != nil {
			return err
		}
		countInt++

		return c.Render(http.StatusOK, "counter", renderer.RenderWidget{Data: countInt})
	})

	// for the main page
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", renderer.RenderPage{Title: "Hi welcome", Layout: "main-layout"})
	})

	e.GET("/about", func(c echo.Context) error {
		return c.Render(http.StatusOK, "about.html", renderer.RenderPage{Title: "About", Layout: "main-layout"})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
