package gobikeindex

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/gorilla/schema"
)

// SearchCloseSerials This endpoint accepts the same parameters as the root /search endpoint.
type SearchCloseSerials struct {
	Bikes Bikes
}

// SearchCloseSerials returns matches that are off of the submitted serial number by less than 3 characters (postgres levenshtein, if you’re curious).
// This endpoint accepts the same parameters as the root `/search` endpoint.
func (bi *BikeIndex) SearchCloseSerials(req SearchReq) (SearchCloseSerials, error) {
	encoder := schema.NewEncoder()
	values := url.Values{}

	// encode variables
	if err := encoder.Encode(req, values); err != nil {
		return SearchCloseSerials{}, err
	}

	// validate values
	if !ValidateStolenness(req.Stolenness) {
		return SearchCloseSerials{}, errInvalidStolennessVal
	}
	if req.Serial == "" {
		return SearchCloseSerials{}, errSerialRequired
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

	url := fmt.Sprintf("%s/search/close_serials?%s", bi.APIURL, values.Encode())
	res := SearchCloseSerials{}
	err := bi.getJSON(url, &res)
	return res, err
}
