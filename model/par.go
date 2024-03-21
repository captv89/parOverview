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
	Reported                bool   `json:"Reported?"`
	ReportedTo              string `json:"Reported to..."`
	CoastalStateActionTaken string `json:"Coastal State Action Taken"`
	MSCCirc                 int    `json:"MSC/Circ"`
}
