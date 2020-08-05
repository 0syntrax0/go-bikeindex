package gobikeindex

// Component bike components
type Component struct {
	ID               int     `json:"id"`
	Description      string  `json:"description"`
	SerialNumber     string  `json:"serial_number"`
	ComponentType    string  `json:"component_type"`
	ComponentGroup   string  `json:"component_group"`
	Rear             *string `json:"rear"`
	Front            *string `json:"front"`
	ManufacturerName *string `json:"manufacturer_name"`
	ModelName        string  `json:"model_name"`
	Year             *string `json:"year"`
}

// Components -
type Components []*Component
