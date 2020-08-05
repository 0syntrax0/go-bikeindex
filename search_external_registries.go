package gobikeindex

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/gorilla/schema"
)

// SearchExternalRegistries This endpoint accepts the same parameters as the root /search endpoint.
type SearchExternalRegistries struct {
	Bikes Bikes
}

// SearchExternalRegistries accepts a serial number and searches external bike registries for it.
// If exact matches are found, only those will be returned.
// If no exact matches are found, partial matches are returned.
func (bi *BikeIndex) SearchExternalRegistries(req SearchReq) (SearchExternalRegistries, error) {
	encoder := schema.NewEncoder()
	values := url.Values{}

	// encode variables
	if err := encoder.Encode(req, values); err != nil {
		return SearchExternalRegistries{}, err
	}

	// validate values
	if !ValidateStolenness(req.Stolenness) {
		return SearchExternalRegistries{}, errInvalidStolennessVal
	}
	if req.Serial == "" {
		return SearchExternalRegistries{}, errSerialRequired
	}

	// remove 0 value elements to prevent unwanted search results
	if req.Page == 0 {
		values.Del("Page")
	}
	if req.PerPage == 0 {
		values.Del("PerPage")
	} else if req.PerPage > 0 {
		values.Del("PerPage")
		values.Add("per_page", strconv.Itoa(req.PerPage))
	}

	url := fmt.Sprintf("%s/search/external_registries?%s", bi.APIURL, values.Encode())
	res := SearchExternalRegistries{}
	err := bi.getJSON(url, &res)
	return res, err
}
