// Sample run-helloworld using Echo
package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Define the handler for the root route
	e.GET("/", func(c echo.Context) error {
		name := os.Getenv("ACCESS")
		if name == "" {
			name = "World"
		}
		return c.String(http.StatusOK, "Hello "+name+"!\n")
	})

	// Determine port for HTTP service
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	e.Logger.Printf("Listening on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
