package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Reter => handenler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
