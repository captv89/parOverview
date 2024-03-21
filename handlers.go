package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/captv89/parOverview/data"
	"github.com/captv89/parOverview/templates"
	"github.com/captv89/parOverview/templates/pages"

	"github.com/angelofallars/htmx-go"
	"github.com/labstack/echo/v4"
)

// indexViewHandler handles a view for the index page.
func indexViewHandler(c echo.Context) error {
	// Set the response content type to HTML.
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Define template meta tags.
	metaTags := pages.MetaTags(
		"gowebly, htmx example page, go with htmx",               // define meta keywords
		"Welcome to example! You're here because it worked out.", // define meta description
	)

	// Define template body content.
	bodyContent := pages.IndexBodyContent(
		"Welcome to example!",                // define h1 text
		"You're here because it worked out.", // define p text
	)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Welcome to example!", // define title text
		metaTags,              // define meta tags
		bodyContent,           // define body content
	)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)
}

// tabularViewHandler handles a view for the index page.
func tabularViewHandler(c echo.Context) error {
	// Set the response content type to HTML.
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Define template meta tags.
	metaTags := pages.MetaTags(
		"gowebly, htmx example page, go with htmx",               // define meta keywords
		"Welcome to example! You're here because it worked out.", // define meta description
	)

	d := data.LoadData("data/sample.json")

	// Define template body content.
	bodyContent := pages.TabularBody(d)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Tabular!",  // define title text
		metaTags,    // define meta tags
		bodyContent, // define body content
	)

	// Sleep for 2 seconds to simulate a slow server response.
	time.Sleep(2 * time.Second)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)
}

// showContentAPIHandler handles an API endpoint to show content.
func showContentAPIHandler(c echo.Context) error {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	// Write HTML content.
	c.Response().Write([]byte("<p>🎉 Yes, <strong>htmx</strong> is ready to use! (<code>GET /api/hello-world</code>)</p>"))

	return htmx.NewResponse().Write(c.Response().Writer)
}
