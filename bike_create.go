package gobikeindex

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

// BikeCreate creates a new bike
type BikeCreate struct {
	BikeEdit
}

// BikeCreate creates a new bike
// Requires `write_bikes` in the access token you use to create the bike.
// IMPORTANT: Ownership: Bikes you create will be created by the user token you authenticate with, but they will be sent to the email address you specify.
func (bi *BikeIndex) BikeCreate(id int64, options BikeCreate) (*BikeCreate, *Exception, error) {
	// check for required fields
	if id == 0 {
		return nil, nil, errBikeIDRequired
	} else if options.Serial == "" {
		return nil, nil, errBikeSerialRequired
	} else if options.ManufacturerName == "" && options.ManufacturerID == "" {
		return nil, nil, errBikeManufacturerRequired
	} else if options.OwnerEmail == nil || *options.OwnerEmail == "" {
		return nil, nil, errBikeOwnerEmailRequired
	} else if options.Color == "" {
		return nil, nil, errBikeColorRequired
	}

	// convert options to HTTP form
	form, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}
	url := bi.buildURL("bikes", strconv.FormatInt(id, 10)) //".json"
	res, err := bi.post(form, url)
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

	bc := new(BikeCreate)
	err = decoder.Decode(bc)
	return bc, nil, err
}
