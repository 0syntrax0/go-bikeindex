package gobikeindex

// StolenRecord info about stolen bikes
type StolenRecord struct {
	Phone                  string `json:"phone"`
	City                   string `json:"city"`
	Country                string `json:"country"`
	Zipcode                string `json:"zipcode"`
	State                  string `json:"state"`
	Address                string `json:"address"`
	DateStolen             string `json:"date_stolen"`
	PoliceReportNumber     string `json:"police_report_number"`
	PoliceReportDepartment string `json:"police_report_department"`
	ShowAddress            bool   `json:"show_address"`
	TheftDescription       string `json:"theft_description"`
}
