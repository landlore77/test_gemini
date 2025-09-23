package requests

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}
