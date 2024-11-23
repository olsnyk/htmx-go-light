package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

type Icecream struct {
	Icecream int
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	icecream := Icecream{Icecream: 0}
	e.Renderer = newTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", icecream)
	})

	e.POST("/icecream", func(c echo.Context) error {
		icecream.Icecream++
		return c.Render(200, "icecream", icecream)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
