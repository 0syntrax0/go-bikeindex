package gobikeindex

import (
	"fmt"
	"net/url"

	"github.com/gorilla/schema"
)

// SearchCountReq params to send
type SearchCountReq struct {
	Serial       string
	Query        string
	Manufacturer string
	Colors       []string
	Location     string
	Distance     string // Distance in miles from location for proximity search
	Stolenness   string
}

// SearchCount Bikeindex API response
// https://bikeindex.org/documentation/api_v3#!/search/GET_version_search_count_format_get_1
type SearchCount struct {
	Proximity int // The count of matching stolen bikes within the proximity of your search.
	Stolen    int
	Non       int
}

// SearchCount sends a request to search for bikes
func (bi *BikeIndex) SearchCount(req SearchCountReq) (SearchCount, error) {
	encoder := schema.NewEncoder()
	values := url.Values{}

	// encode variables
	if err := encoder.Encode(req, values); err != nil {
		return SearchCount{}, err
	}

	// validate values
	if !ValidateStolenness(req.Stolenness) {
		return SearchCount{}, errInvalidStolennessVal
	}

	url := fmt.Sprintf("%s/search/count?%s", bi.APIURL, values.Encode())
	res := SearchCount{}
	err := bi.getJSON(url, &res)
	return res, err
}
