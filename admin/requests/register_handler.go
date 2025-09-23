package requests

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}
