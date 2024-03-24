package main

import (
	"encoding/json"
	"fmt"
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
		"piracy, robbery, maritime, hostage, ship, vessel, armed, attack, overview",
		"The piracy and armed robbery overview sourced from the IMB Piracy Reporting Center.",
	)

	// Define template body content.
	mapComponent := pages.MapComponent(data.GeoParData)
	bodyContent := pages.IndexBodyContent(mapComponent)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Maritime Piracy & Robbery Overview", // define title text
		metaTags,                             // define meta tags
		bodyContent,
	)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)
}

// tabularViewHandler handles a view for the tabular page.
func tabularViewHandler(c echo.Context) error {
	// Set the response content type to HTML.
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Define template meta tags.
	metaTags := pages.MetaTags(
		"dataset, piracy, robbery, maritime, hostage, ship, vessel, armed, attack, overview",
		"The piracy and armed robbery overview dataset of maritime incidents.",
	)

	// Define template body content.
	bodyContent := pages.TabularBody(data.GeoParData)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Maritime Piracy & Robbery Dataset", // define title text
		metaTags,                            // define meta tags
		bodyContent,
	)

	// Sleep for 2 seconds to simulate a slow server response.
	time.Sleep(2 * time.Second)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)
}

// mapViewHandler handles a view for the map page.
func mapViewHandler(c echo.Context) error {
	// Set the response content type to HTML.
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Define template meta tags.
	metaTags := pages.MetaTags(
		"map, maritime piracy, armed robbery, overview",
		"The piracy and armed robbery overview map of maritime incidents.",
	)

	// Define template body content.
	bodyContent := pages.MapComponent(data.GeoParData)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Maritime Piracy & Robbery Overview", // define title text
		metaTags,                             // define meta tags
		bodyContent,                          // define body content
	)

	// Sleep for 2 seconds to simulate a slow server response.
	time.Sleep(2 * time.Second)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)
}

// chartViewHandler handles a view for the map page.
func chartViewHandler(c echo.Context) error {
	// Set the response content type to HTML.
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	// Define template meta tags.
	metaTags := pages.MetaTags(
		"chart, maritime piracy, armed robbery, overview, trends",
		"The piracy and armed robbery trends over time of maritime incidents.",
	)

	// Data for the chart
	incidentByYear := data.IncidentByYear(data.ParData)

	// Define template body content.
	chart := pages.ChartComponent("Incidents by Year", incidentByYear)

	bodyContent := pages.ChartBody(chart)

	// Define template layout for index page.
	indexTemplate := templates.Layout(
		"Maritime Piracy & Robbery Trend", // define title text
		metaTags,                          // define meta tags
		bodyContent,                       // define body content
	)

	// Sleep for 2 seconds to simulate a slow server response.
	time.Sleep(2 * time.Second)

	return htmx.NewResponse().RenderTempl(c.Request().Context(), c.Response().Writer, indexTemplate)
}

// API handlers
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

// incidentsByParamsHandler handles an API endpoint to get incidents by parameters.
func incidentsByParamsHandler(c echo.Context) error {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(c.Request()) {
		// If not, return HTTP 400 error.
		c.Response().WriteHeader(http.StatusBadRequest)
		slog.Error("request API", "method", c.Request().Method, "status", http.StatusBadRequest, "path", c.Request().URL.Path)
		return echo.NewHTTPError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	// Get query parameters from the request.
	by := c.QueryParam("by")

	// Define a response message.
	var title string
	var cdata interface{}

	switch by {
	case "year":
		// Data for the chart
		cdata = data.IncidentByYear(data.ParData)
		title = "Incidents by Year"
	case "month":
		// Data for the chart
		cdata = data.IncidentByMonthYear(data.ParData)
		title = "Incidents by Month"
	case "area":
		// Data for the chart
		cdata = data.IncidentByArea(data.ParData)
		title = "Incidents by Area"
	case "shipType":
		// Data for the chart
		cdata = data.IncidentByShipType(data.ParData)
		title = "Incidents by Ship Type"
	default:
		// Data for the chart
		cdata = data.IncidentByYear(data.ParData)
		title = "Incidents by Year"
	}

	// JSON the response data.
	resData := map[string]interface{}{
		"title": title,
		"data":  cdata,
	}

	resEvent := map[string]interface{}{
		"updateChartData": resData,
	}

	// Convert the response map to JSON
	resJSONData, err := json.Marshal(resEvent)
	if err != nil {
		fmt.Println("Error:", err)
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	// Convert the JSON bytes to a string
	resJSONString := string(resJSONData)

	// Write JSON content.
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().Writer.Header().Set("HX-Trigger", resJSONString)
	// Write JSON content.
	return c.String(http.StatusOK, "")
}
