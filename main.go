package main

import (
	"html/template"
	"io"
	"net/http"
	"test2/router"
	"test2/utils/config"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	config.LoadConfig("config/config.yaml")

	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// Login check middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Allow access to / and /login without authentication
			if c.Path() == "/" || c.Path() == "/login" || c.Path() == "/actions/login" || c.Path() == "/register" || c.Path() == "/actions/register" {
				return next(c)
			}

			sess, _ := session.Get("session", c)
			if auth, ok := sess.Values["authenticated"].(bool); !ok || !auth {
				return c.Redirect(http.StatusFound, "/login")
			}
			return next(c)
		}
	})

	t := &Template{
		templates: template.Must(template.ParseGlob("pages/*.html")),
	}
	e.Renderer = t

	router.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
