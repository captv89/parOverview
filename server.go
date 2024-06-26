package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	gowebly "github.com/gowebly/helpers"
)

// runServer runs a new HTTP server with the loaded environment variables.
func runServer() error {
	// Validate environment variables.
	port, err := strconv.Atoi(gowebly.Getenv("BACKEND_PORT", "3000"))
	if err != nil {
		return err
	}
	host := gowebly.Getenv("BACKEND_HOST", "localhost")

	// Create a new Echo server.
	e := echo.New()

	// Add Echo middlewares.
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.CORS())
	e.Use(middleware.CSRF())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	e.Use(middleware.Recover())

	// Handle static files.
	e.Static("/static", "./static")

	// Handle index page view.
	e.GET("/", indexViewHandler)

	// Handle table page view.
	e.GET("/table", tabularViewHandler)

	// Handle map page view.
	e.GET("/map", mapViewHandler)

	// Handle chart page view.
	e.GET("/chart", chartViewHandler)

	// Handle API endpoints.

	// Incidents API
	e.GET("/api/incidents", incidentsByParamsHandler)

	// Create a new server instance with options from environment variables.
	// For more information, see https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/
	server := http.Server{
		Addr:         fmt.Sprintf("%v:%d", host, port),
		Handler:      e, // handle all Echo routes
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Send log message.
	slog.Info("Starting server...", "port", port)

	return server.ListenAndServe()
}
