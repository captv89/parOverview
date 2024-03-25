package model

type ParReport struct {
	Date                    string `json:"Date"`
	ShipName                string `json:"Ship Name"`
	ShipType                string `json:"Ship Type"`
	IMONo                   string `json:"IMO No.,omitempty"`
	Area                    string `json:"Area"`
	Latitude                string `json:"Latitude"`
	Longitude               string `json:"Longitude"`
	IncidentDetails         string `json:"Incident details"`
	ConsequencesForCrewEtc  string `json:"Consequences for crew etc"`
	ActionTakenByMasterCrew string `json:"Action taken by master/crew"`
	Reported                string `json:"Reported?"`
	ReportedTo              string `json:"Reported to..."`
	CoastalStateActionTaken string `json:"Coastal State Action Taken"`
	MSCCirc                 string `json:"MSC/Circ"`
}

type CleanParReport struct {
	Date                    string `json:"Date"`
	ShipName                string `json:"ShipName"`
	ShipType                string `json:"ShipType"`
	IMONo                   string `json:"IMO"`
	Area                    string `json:"Area"`
	Latitude                string `json:"Latitude"`
	Longitude               string `json:"Longitude"`
	IncidentDetails         string `json:"IncidentDetails"`
	ConsequencesForCrewEtc  string `json:"Consequences"`
	ActionTakenByMasterCrew string `json:"ActionTakenByShip"`
	Reported                string `json:"Reported"`
	ReportedTo              string `json:"ReportedTo"`
	CoastalStateActionTaken string `json:"CoastalStateActionTaken"`
	MSCCirc                 string `json:"MSCirc"`
}

type GeoParReport struct {
	Date                    string  `json:"Date"`
	ShipName                string  `json:"ShipName"`
	ShipType                string  `json:"ShipType"`
	IMONo                   string  `json:"IMO"`
	Area                    string  `json:"Area"`
	Latitude                float64 `json:"Latitude"`
	Longitude               float64 `json:"Longitude"`
	IncidentDetails         string  `json:"IncidentDetails"`
	ConsequencesForCrewEtc  string  `json:"Consequences"`
	ActionTakenByMasterCrew string  `json:"ActionTakenByShip"`
	Reported                string  `json:"Reported"`
	ReportedTo              string  `json:"ReportedTo"`
	CoastalStateActionTaken string  `json:"CoastalStateActionTaken"`
	MSCCirc                 string  `json:"MSCirc"`
}
