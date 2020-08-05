package gobikeindex

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

// BikeUpdate params to update
type BikeUpdate struct {
	// optional
	RearWheelBSD        int    `url:"RearWheelBSD,omitempty" json:"rear_wheel_bsd,omitempty"`     // Rear wheel iso bsd (has to be one of the selections)
	RearTireNarrow      bool   `url:"RearTireNarrow,omitempty" json:"rear_tire_narrow,omitempty"` // Is it a skinny tire?
	FrontWheelBSD       string `url:",omitempty" json:"front_wheel_bsd,omitempty"`                // Copies rear_wheel_bsd if not set
	FrontTireNarrow     bool   `url:"FrontWheelBSD,omitempty" json:"front_tire_narrow,omitempty"` // Copies rear_tire_narrow if not set
	FrameModel          string `url:"FrameModel,omitempty" json:"frame_model,omitempty"`          // What frame model?
	Year                int    `url:"Year,omitempty" json:"year,omitempty"`
	Description         string `url:"Description,omitempty" json:"description,omitempty"`
	PrimaryFrameColor   string `url:"PrimaryFrameColor,omitempty" json:"primary_frame_color,omitempty"`
	SecondaryFrameColor string `url:"SecondaryFrameColor,omitempty" json:"secondary_frame_color,omitempty"`
	TertiaryFrameColor  string `url:"TertiaryFrameColor,omitempty" json:"tertiary_frame_color,omitempty"`
	RearGearTypeSlug    string `url:"RearGearTypeSlug,omitempty" json:"rear_gear_type_slug,omitempty"`   // rear gears (has to be one of the `selections`)
	FrontGearTypeSlug   string `url:"FrontGearTypeSlug,omitempty" json:"front_gear_type_slug,omitempty"` // front gears (has to be one of the `selections`)
	HandleBarTypeSlug   string `url:"HandleBarTypeSlug,omitempty" json:"handlebar_type_slug,omitempty"`  // handlebar type (has to be one of the `selections`)
	NoNotify            bool   `url:"NoNotify,omitempty" json:"no_notify,omitempty"`                     // On create or ownership change, don’t notify the new owner.
	IsForSale           bool   `url:"IsForSale,omitempty" json:"is_for_sale,omitempty"`
	FrameMaterial       string `url:"FrameMaterial,omitempty" json:"frame_material,omitempty"`
}

//https://bikeindex.org/documentation/api_v3#!/bikes/PUT_version_bikes_id_format_put_1

// BikeUpdate update a bike OWNED by the access token
// Note: Access to this endpoint is only available to select api clients
// Requires `read_user` in the access token you use to send the notification.
func (bi *BikeIndex) BikeUpdate(id int64, options BikeUpdate) (*BikeUpdate, *Exception, error) {
	if id == 0 {
		return nil, nil, errBikeIDRequired
	}

	// convert options to HTTP form
	form, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	url := bi.buildURL("bikes", strconv.FormatInt(id, 10)) //".json"
	res, err := bi.put(form, url)
	if err != nil {
		return nil, nil, err
	}

	decoder := json.NewDecoder(res.Body)

	// handle NULL response
	if res.StatusCode != http.StatusOK {
		exception := new(Exception)
		err = decoder.Decode(exception)
		return nil, exception, err
	}

	bur := new(BikeUpdate)
	err = decoder.Decode(bur)
	return bur, nil, err
}
