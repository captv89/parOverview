package data

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"regexp"
	"strconv"

	"github.com/captv89/parOverview/model"
)

var (
	ParData    []model.ParReport
	GeoParData []model.GeoParReport
)

func parseCoordinates(coord string) float64 {
	re := regexp.MustCompile(`(\d+)°\s*(\d+\.?\d*)'`)
	matches := re.FindAllStringSubmatch(coord, -1)
	if len(matches) == 0 {
		return 0
	}

	degrees, _ := strconv.ParseFloat(matches[0][1], 64)
	minutes, _ := strconv.ParseFloat(matches[0][2], 64)

	// Convert to decimal
	degrees += minutes / 60

	// Check if S or W
	if coord[len(coord)-1] == 'S' || coord[len(coord)-1] == 'W' {
		degrees = -degrees
	}

	return degrees
}

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

// LoadGeoData loads data from a file.
func LoadGeoData(data []model.ParReport) []model.GeoParReport {
	var reports []model.GeoParReport

	for i := 0; i < len(data); i++ {
		report := data[i]
		lat := parseCoordinates(report.Latitude)
		lon := parseCoordinates(report.Longitude)

		reports = append(reports, model.GeoParReport{
			Date:                    report.Date,
			ShipName:                report.ShipName,
			ShipType:                report.ShipType,
			IMONo:                   report.IMONo,
			Area:                    report.Area,
			Latitude:                lat,
			Longitude:               lon,
			IncidentDetails:         report.IncidentDetails,
			ConsequencesForCrewEtc:  report.ConsequencesForCrewEtc,
			ActionTakenByMasterCrew: report.ActionTakenByMasterCrew,
			Reported:                report.Reported,
			ReportedTo:              report.ReportedTo,
			CoastalStateActionTaken: report.CoastalStateActionTaken,
			MSCCirc:                 report.MSCCirc,
		})
	}

	return reports
}
