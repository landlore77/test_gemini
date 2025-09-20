package main

import (
	"log"
	"test1/config"
	"test1/router"
	"test1/utils"
	"test1/request" // New import

	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Main function started.")
	config.LoadConfig()
	log.Println("Config loaded.")

	e := echo.New()
	log.Println("Echo instance created.")

	// Initialize the template renderer
	// e.Renderer = utils.NewTemplateRenderer("pages/*.html")
	renderer := utils.NewTemplateRenderer("src/pages/*.html")
	e.Renderer = renderer
	log.Println("Template renderer initialized.")

	for _, route := range router.GetRoutes() {
		e.GET(route.Path, route.Handler)
	}
	log.Println("Routes registered.")
	// Handle POST for /actions/register
	e.POST("/actions/register", request.RegisterActionHandler)
	log.Println("Register action handler registered.")

	log.Println("Starting server on :8080...")
	e.Logger.Fatal(e.Start(":8080"))
}