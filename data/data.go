package data

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"github.com/captv89/parOverview/model"
)

// LoadData loads data from a file.
func LoadData(filePath string) []model.ParReport {
	var reports []model.ParReport

	// Read the json file
	jsonFile, err := os.Open(filePath)
	if err != nil {
		slog.Error("Failed to open file", "details", err.Error())
	}

	defer jsonFile.Close()

	// Read the json file
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		slog.Error("Failed to read file", "details", err.Error())
	}

	// Unmarshal the json file
	err = json.Unmarshal(byteValue, &reports)
	if err != nil {
		slog.Error("Failed to unmarshal file", "details", err.Error())
	}

	return reports
}
