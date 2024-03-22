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

// IncidentByYear returns the number of incidents by year.
func IncidentByYear(data []model.ParReport) map[string]int {
	incidents := make(map[string]int)

	for i := 0; i < len(data); i++ {
		report := data[i]
		year := report.Date[:4]

		// Check if the year is already in the map
		if _, ok := incidents[year]; ok {
			incidents[year]++
		} else {
			incidents[year] = 1
		}
	}

	return incidents
}

// IncidentByArea returns the number of incidents by area.
func IncidentByArea(data []model.ParReport) map[string]int {
	incidents := make(map[string]int)

	for i := 0; i < len(data); i++ {
		report := data[i]

		// Check if the area is already in the map
		if _, ok := incidents[report.Area]; ok {
			incidents[report.Area]++
		} else {
			incidents[report.Area] = 1
		}
	}

	return incidents
}

// IncidentByShipType returns the number of incidents by ship type.
func IncidentByShipType(data []model.ParReport) map[string]int {
	incidents := make(map[string]int)

	for i := 0; i < len(data); i++ {
		report := data[i]

		// Check if the ship type is already in the map
		if _, ok := incidents[report.ShipType]; ok {
			incidents[report.ShipType]++
		} else {
			incidents[report.ShipType] = 1
		}
	}

	return incidents
}

// IncidentByMonth returns the number of incidents by month.
func IncidentByMonth(data []model.ParReport) map[string]int {
	incidents := make(map[string]int)

	for i := 0; i < len(data); i++ {
		report := data[i]
		month := report.Date[5:7]
		month = GetMonthNames(month)

		// Check if the month is already in the map
		if _, ok := incidents[month]; ok {
			incidents[month]++
		} else {
			incidents[month] = 1
		}
	}

	return incidents
}

// IncidentByMonthYear returns the number of incidents by month and year.
func IncidentByMonthYear(data []model.ParReport) map[string]map[string]int {
	incidents := make(map[string]map[string]int)

	for i := 0; i < len(data); i++ {
		report := data[i]
		month := report.Date[5:7]
		month = GetMonthNames(month)
		year := report.Date[:4]

		// Check if the month is already in the map
		if _, ok := incidents[year]; ok {
			if _, ok := incidents[year][month]; ok {
				incidents[year][month]++
			} else {
				incidents[year][month] = 1
			}
		} else {
			incidents[year] = make(map[string]int)
			incidents[year][month] = 1
		}
	}

	return incidents
}

// GetMonthNames returns the month names.
func GetMonthNames(num string) string {
	months := map[string]string{
		"01": "January",
		"02": "February",
		"03": "March",
		"04": "April",
		"05": "May",
		"06": "June",
		"07": "July",
		"08": "August",
		"09": "September",
		"10": "October",
		"11": "November",
		"12": "December",
	}

	return months[num]
}
