package main

import (
	"html/template"
	"io"
	"test2/router"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	t := &Template{
		templates: template.Must(template.ParseGlob("pages/*.html")),
	}
	e.Renderer = t

	router.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
