package main

import (
	"fmt"
    "os"
    // "strings"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load variables
	// Using var keyword
    var PORT string
    PORT = os.Getenv("PORT")
	fmt.Println("PORT:", os.Getenv("PORT"))

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	// Start server
	e.Logger.Fatal(e.Start(":"+PORT))
}
