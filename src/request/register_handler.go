package request

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RegisterPageHandler renders the register page
func RegisterPageHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}
