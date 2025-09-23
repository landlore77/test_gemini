package router

import (
	"test2/admin/requests"

	"github.com/labstack/echo/v4"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func InitRoutes(e *echo.Echo) {
	routes := []Route{
		{Method: "GET", Path: "/", Handler: requests.HomeHandler},
		{Method: "GET", Path: "/login", Handler: requests.LoginHandler},
		{Method: "GET", Path: "/register", Handler: requests.RegisterHandler},
		{Method: "GET", Path: "/admin_list", Handler: requests.AdminListHandler},
		{Method: "POST", Path: "/actions/register", Handler: requests.RegisterActionHandler},
	}

	for _, route := range routes {
		switch route.Method {
		case "GET":
			e.GET(route.Path, route.Handler)
		case "POST":
			e.POST(route.Path, route.Handler)
		}
	}
}
