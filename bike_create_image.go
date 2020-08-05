package gobikeindex

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

// BikeNewImage adds new image to a bike
type BikeNewImage struct {
	ImagePath string
}

// BikeNewImage adds new image to a bike
func (bi *BikeIndex) BikeNewImage(id int64, options BikeNewImage) (*BikeNewImage, *Exception, error) {
	// convert options to HTTP form
	form, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}
	url := bi.buildURL("bikes", strconv.FormatInt(id, 10), "image") //".json"
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

	bni := new(BikeNewImage)
	err = decoder.Decode(bni)
	return bni, nil, err
}
