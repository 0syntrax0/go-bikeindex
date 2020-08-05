package gobikeindex

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

// BikeUpdate params to update
type BikeUpdate struct {
	BikeEdit
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
