package gobikeindex

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/gorilla/schema"
)

// SearchSerialsContaining This endpoint accepts the same parameters as the root /search endpoint.
type SearchSerialsContaining struct {
	Bikes Bikes
}

// SearchSerialsContaining It returns bikes with partially-matching serial numbers (for which the requested serial is a substring).
// This endpoint accepts the same parameters as the root `/search` endpoint.
func (bi *BikeIndex) SearchSerialsContaining(req SearchReq) (SearchSerialsContaining, error) {
	encoder := schema.NewEncoder()
	values := url.Values{}

	// encode variables
	if err := encoder.Encode(req, values); err != nil {
		return SearchSerialsContaining{}, err
	}

	// validate values
	if !ValidateStolenness(req.Stolenness) {
		return SearchSerialsContaining{}, errInvalidStolennessVal
	}
	if req.Serial == "" {
		return SearchSerialsContaining{}, errSerialRequired
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

	url := fmt.Sprintf("%s/search/serials_containing?%s", bi.APIURL, values.Encode())
	res := SearchSerialsContaining{}
	err := bi.getJSON(url, &res)
	return res, err
}
