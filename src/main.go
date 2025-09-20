package main

import (
	"log"
	"test1/config"
	"test1/router"
	"test1/utils"

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
		switch route.Method {
		case "GET":
			e.GET(route.Path, route.Handler)
		case "POST":
			e.POST(route.Path, route.Handler)
		// Add other methods as needed
		}
	}
	log.Println("Routes registered.")
	// Removed direct POST registration as it's now handled by router.GetRoutes()
	log.Println("Register action handler registered.")

	log.Println("Starting server on :8080...")
	e.Logger.Fatal(e.Start(":8080"))
}