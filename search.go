package gobikeindex

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/gorilla/schema"
)

var (
	// errors
	errInvalidStolennessVal = errors.New("invalid stolenness option")
	errSerialRequired       = errors.New("serial is required")
)

// SearchReq - parameters to send
type SearchReq struct {
	Page         int      `json:"page,omitempty"`     // page of results to fetch
	PerPage      int      `json:"per_page,omitempty"` // bikes per page (max 100)
	Serial       string   `json:"serial,omitempty"`   // serial, homoglyph matched
	Query        string   `json:"query,omitempty"`    // full text search
	Manufacturer string   `json:"manufacturer,omitempty"`
	Colors       []string `json:"colors,omitempty"`
	Location     string   `json:"location,omitempty"` // I.P. or Location address for proximity search - location is ignored unless stolenness is “proximity”
	Distance     string   `json:"distance,omitempty"` // Distance in miles from location for proximity search
	Stolenness   string   `json:"stolenness,omitempty"`
}

// Search Bikeindex API response
// https://bikeindex.org/documentation/api_v3#!/search/GET_version_search_format_get_0
type Search struct {
	Bikes Bikes
}

// Search sends a request to search for bikes
func (bi *BikeIndex) Search(req SearchReq) (Search, error) {
	encoder := schema.NewEncoder()
	values := url.Values{}

	// encode variables
	if err := encoder.Encode(req, values); err != nil {
		return Search{}, err
	}

	// validate values
	if !ValidateStolenness(req.Stolenness) {
		return Search{}, errInvalidStolennessVal
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

	// url := fmt.Sprintf("%s/search?%s", bi.APIURL, values.Encode())
	res := Search{}
	err := bi.getJSON(bi.buildURL("/search?", values.Encode()), &res)
	return res, err
}
