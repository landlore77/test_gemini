package request

import (
	"net/http"
	"test1/config"

	"github.com/labstack/echo/v4"
)

// LoginHandler handles the login request
func LoginHandler(c echo.Context) error {
	data := map[string]interface{}{
		"googleClientID": config.Cfg.GoogleClientID,
	}
	return c.Render(http.StatusOK, "login.html", data)
}
