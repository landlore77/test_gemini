package requests

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func LogoutActionHandler(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Values["authenticated"] = false
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, "/login")
}
