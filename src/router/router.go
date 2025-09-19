package router

import (
	"test1/request"
	"test1/admin"

	"github.com/labstack/echo/v4"
)

// Route defines a route for the web server
type Route struct {
	Path    string
	Handler echo.HandlerFunc
}

// routing data
var adminListHandler = admin.AdminListHandler // Explicitly use the handler

var routingData = []Route{
	{"/", request.HelloHandler},
	{"/login", request.LoginHandler},
	{"/register", request.RegisterPageHandler},
	{"/admin_list", adminListHandler},
	{"/actions/register", request.RegisterActionHandler},
}

func GetRoutes() []Route {
	return routingData
}
