package main

import (
	"test1/config"
	"test1/router"
	"test1/utils"
	"test1/request" // New import

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()

	e := echo.New()

	// Initialize the template renderer
	// e.Renderer = utils.NewTemplateRenderer("pages/*.html")
	renderer := utils.NewTemplateRenderer("pages/*.html")
	e.Renderer = renderer

	for _, route := range router.GetRoutes() {
		e.GET(route.Path, route.Handler)
	}
	// Handle POST for /actions/register
	e.POST("/actions/register", request.RegisterActionHandler)

	e.Logger.Fatal(e.Start(":8080"))
}