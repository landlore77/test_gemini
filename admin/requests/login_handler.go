package requests

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	return c.String(http.StatusOK, "login page")
}
