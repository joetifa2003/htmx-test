package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/joetifa2003/htmx-test/renderer"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Renderer = renderer.NewRenderer("main-layout") // sets the renderer, which executes templates

	// on dev mode, templates are reparsed on every request (hot reloading)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			e.Renderer = renderer.NewRenderer("main-layout")
			next(c)
			return nil
		}
	})

	e.StaticFS("/", renderer.Assets)

	// for the main page
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", renderer.RenderPage{Title: "Hi welcome"})
	})

	// for the counter widget /renderer/assets/widgets/counter.html
	e.POST("/counter", func(c echo.Context) error {
		count := c.FormValue("count")
		countInt, err := strconv.Atoi(count)
		if err != nil {
			return err
		}
		countInt++

		time.Sleep(1 * time.Second)

		return c.Render(http.StatusOK, "counter", renderer.RenderWidget{Data: countInt})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
