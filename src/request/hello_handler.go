package request

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HelloHandler handles the hello request
func HelloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
