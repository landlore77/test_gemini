package main

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    routes := map[string]echo.HandlerFunc{
        "/": func(c echo.Context) error {
            return c.String(http.StatusOK, "hello")
        },
        "/login": func(c echo.Context) error {
            return c.String(http.StatusOK, "login page")
        },
    }

    for path, handler := range routes {
        e.GET(path, handler)
    }

    e.Logger.Fatal(e.Start(":8080"))
}
