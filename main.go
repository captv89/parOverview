package main

import (
	"log/slog"
	"os"

	"github.com/captv89/parOverview/data"
)

// Main function.
func main() {
	// Load data
	jsonFilePath := os.Getenv("DATA_FILE_PATH")
	data.ParData = data.LoadData(jsonFilePath)
	data.GeoParData = data.LoadGeoData(data.ParData)

	// Run your server.
	if err := runServer(); err != nil {
		slog.Error("Failed to start server!", "details", err.Error())
		os.Exit(1)
	}
}
