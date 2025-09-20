package router

import (
	"test1/request"
	"test1/admin"

	"github.com/labstack/echo/v4"
)

// Route defines a route for the web server
type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

// routing data
var adminListHandler = admin.AdminListHandler // Explicitly use the handler

var routingData = []Route{
	{"GET", "/", request.HelloHandler},
	{"GET", "/login", request.LoginHandler},
	{"GET", "/register", request.RegisterPageHandler},
	{"GET", "/admin_list", adminListHandler},
	{"POST", "/actions/register", request.RegisterActionHandler},
}

func GetRoutes() []Route {
	return routingData
}
