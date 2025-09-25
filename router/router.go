package router

import (
	"test2/admin/requests"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", requests.HomeHandler)
	e.GET("/login", requests.LoginHandler)
	e.GET("/register", requests.RegisterHandler)
	e.GET("/admin_list", requests.AdminListHandler)
	e.POST("/actions/register", requests.RegisterActionHandler)
	e.POST("/actions/login", requests.LoginActionHandler)
	e.POST("/actions/logout", requests.LogoutActionHandler)

	e.GET("/event/pass", requests.EventPassHandler)
	e.POST("/actions/event_pass_add", requests.EventPassAddActionHandler)
	e.GET("/actions/event_pass_delete", requests.EventPassDeleteActionHandler)
}
