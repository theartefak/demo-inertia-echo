package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/theartefak/inertia-echo"
)

func main() {
    e := echo.New()
    e.Use(inertia.Middleware(e))
    e.Static("/", "./public")

    e.GET("/", hello).Name = "welcome"
    e.GET("halo", hello).Name = "halo"

    e.Logger.Fatal(e.Start("127.0.1.1:3000"))
}

func hello(c echo.Context) error {
    // Status, Component Name, Data to pass on
    return c.Render(http.StatusOK, "Welcome", map[string]interface{}{})
}
