package gobikeindex

import (
	"errors"
	"fmt"
)

var (
	// errors
	errBikeIDRequired           = errors.New("bike id required")
	errBikeSerialRequired       = errors.New("bike serial number is required")
	errBikeManufacturerRequired = errors.New("bike manufacturer is required")
	errBikeOwnerEmailRequired   = errors.New("bike owner email is required")
	errBikeColorRequired        = errors.New("bike color is required")
)

// Bikes multiple bikes, duh
type Bikes []Bike

// Bike struct returned from bikeindex
type Bike struct {
	DateStolen             int64    `json:"date_stolen"` // in miliseconds
	Description            *string  `json:"description"`
	FrameColors            []string `json:"frame_colors"`
	FrameModel             *string  `json:"frame_model"`
	ID                     int64    `json:"id"`
	IsStockImg             bool     `json:"is_stock_img"`
	LargeImg               string   `json:"large_img"`
	LocationFound          *string  `json:"location_found"`
	ManufacturerName       string   `json:"manufacturer_name"`
	ExternalID             *int     `json:"external_id"`
	RegistryName           *string  `json:"registry_name"`
	RegistryURL            *string  `json:"registry_url"`
	Serial                 string   `json:"serial"`
	Status                 *string  `json:"status"`
	Stolen                 bool     `json:"stolen"`
	StolenLocation         string   `json:"stolen_location"`
	Thumb                  string   `json:"thumb"`
	Title                  string   `json:"title"`
	URL                    string   `json:"url"`
	Year                   *int     `json:"year"`
	RegistrationCreatedAt  int      `json:"registration_created_at"`
	RegistrationUpdatedAt  int      `json:"registration_updated_at"`
	APIURL                 string   `json:"api_url"`
	ManufacturerID         int      `json:"manufacturer_id"`
	PaintDescription       *string  `json:"paint_description"`
	Name                   string   `json:"name"`
	FrameSize              string   `json:"frame_size"`
	RearTireNarrow         bool     `json:"rear_tire_narrow"`
	FrontTireNarrow        bool     `json:"front_tire_narrow"`
	TypeOfCycle            string   `json:"type_of_cycle"`
	CycleTypeName          *string  `json:"cycle_type_name"`
	TestBike               bool     `json:"test_bike"`
	RearWheelSizeIsoBsd    int      `json:"rear_wheel_size_iso_bsd"`
	FrontWheelSizeIsoBsd   int      `json:"front_wheel_size_iso_bsd"`
	FrameMaterialSlug      string   `json:"frame_material_slug"`
	FrontGearTypeSlug      string   `json:"front_gear_type_slug"`
	RearGearTypeSlug       string   `json:"rear_gear_type_slug"`
	AdditionalRegistration *string  `json:"additional_registration"`
	StolenRecord           *string  `json:"stolen_record"`

	//
	PublicImages *PublicImages `json:"public_images"`
	Components   *Components   `json:"components"`
}

// BikeByID fetch bike info by ID
// https://bikeindex.org/documentation/api_v3#!/bikes/GET_version_bikes_id_format_get_0
func (bi *BikeIndex) BikeByID(id int) (Bike, error) {
	url := fmt.Sprintf("%s/bikes?%d", bi.APIURL, id)
	res := Bike{}
	err := bi.getJSON(url, &res)
	return res, err
}
