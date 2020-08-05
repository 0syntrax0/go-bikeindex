package gobikeindex

// TODO: COMPLETE `BIKEEdit` STRUCT
// BikeEdit fields used when updating/creating a bike
type BikeEdit struct {
	// optional
	Test                bool    `url:"Test,omitempty" json:"test,omitempty"`                         // Enable to test your API, will send data but NOT update
	RearWheelBSD        int     `url:"RearWheelBSD,omitempty" json:"rear_wheel_bsd,omitempty"`       // Rear wheel iso bsd (has to be one of the selections)
	RearTireNarrow      bool    `url:"RearTireNarrow,omitempty" json:"rear_tire_narrow,omitempty"`   // Is it a skinny tire?
	FrontWheelBSD       string  `url:"FrontWheelBSD,omitempty" json:"front_wheel_bsd,omitempty"`     // Copies rear_wheel_bsd if not set
	FrontTireNarrow     bool    `url:"FrontTireNarrow,omitempty" json:"front_tire_narrow,omitempty"` // Copies rear_tire_narrow if not set
	FrameModel          string  `url:"FrameModel,omitempty" json:"frame_model,omitempty"`            // What frame model?
	Year                int     `url:"Year,omitempty" json:"year,omitempty"`
	Description         string  `url:"Description,omitempty" json:"description,omitempty"`
	PrimaryFrameColor   string  `url:"PrimaryFrameColor,omitempty" json:"primary_frame_color,omitempty"`
	SecondaryFrameColor string  `url:"SecondaryFrameColor,omitempty" json:"secondary_frame_color,omitempty"`
	TertiaryFrameColor  string  `url:"TertiaryFrameColor,omitempty" json:"tertiary_frame_color,omitempty"`
	RearGearTypeSlug    string  `url:"RearGearTypeSlug,omitempty" json:"rear_gear_type_slug,omitempty"`   // rear gears (has to be one of the `selections`)
	FrontGearTypeSlug   string  `url:"FrontGearTypeSlug,omitempty" json:"front_gear_type_slug,omitempty"` // front gears (has to be one of the `selections`)
	HandleBarTypeSlug   string  `url:"HandleBarTypeSlug,omitempty" json:"handlebar_type_slug,omitempty"`  // handlebar type (has to be one of the `selections`)
	NoNotify            bool    `url:"NoNotify,omitempty" json:"no_notify,omitempty"`                     // On create or ownership change, don’t notify the new owner.
	IsForSale           bool    `url:"IsForSale,omitempty" json:"is_for_sale,omitempty"`
	FrameMaterial       string  `url:"FrameMaterial,omitempty" json:"frame_material,omitempty"`
	OrganizationSlug    string  `url:"OrganizationSlug,omitempty" json:"organization_slug,omitempty"`
	Color               string  `url:"Color" json:"color"`
	OwnerEmail          *string `url:"OwnerEmail,omitempty" json:"owner_email,omitempty"`
	ManufacturerName    string  `url:"ManufacturerName,omitempty" json:"manufacturer,omitempty"`
	ManufacturerID      string  `url:"ManufacturerID,omitempty" json:"manufacturer_id,omitempty"`
	Serial              string  `url:"Serial,omitempty" json:"serial,omitempty"`
}
