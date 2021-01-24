package gobikeindex

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

var (
	// OrganizationListKind list of the kind of organization
	OrganizationListKind = []string{
		"bike_shop",
		"bike_advocacy",
		"bike_manufacturer",
		"law_enforcement",
		"other",
		"property_management",
		"school",
		"software",
	}

	// errors
	errOrganizationKind = errors.New("invalid organization kind")
)

// Organization - Note: Access to this endpoint is only available to select api clients.
// OrganizationReq - Operations about organizations
type Organization struct {
	Name      string                 `json:"name"` // The organization name
	Website   string                 `json:"website"`
	Kind      string                 `json:"kind"`
	Locations []OrganizationLocation `json:"locations"`
}

type OrganizationLocation struct {
	Name    string  `json:"name"` // The location’s name
	Phone   *string `json:"phone,omitempty"`
	Street  string  `json:"street"`
	City    string  `json:"city"`
	State   *string `json:"state,omitempty"`
	Zipcode *string `json:"zipcode,omitempty"`
	Country *string `json:"country,omitempty"`
}

// OrganizationsCreate - creates a new organization
// Requires `write_organizations` in the access token you use to create the organization
func (bi *BikeIndex) OrganizationsCreate(o *Organization) (*Organization, *Exception, error) {
	// validate kind
	if o.Kind != "" && !bi.IsValidOrganizationType(o.Kind) {
		return nil, nil, errOrganizationKind
	}
	// convert o to HTTP form
	form, err := query.Values(o)
	if err != nil {
		return nil, nil, err
	}
	// create url and make call
	res, err := bi.post(form, bi.buildURL("organizations"))
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

	org := new(Organization)
	err = decoder.Decode(org)
	return org, nil, err
}

// IsValidOrganizationType checks if a given organization type exists
func (bi *BikeIndex) IsValidOrganizationType(s string) bool {
	s = strings.ToLower(s)
	for _, t := range OrganizationListKind {
		if t == s {
			return true
		}
	}
	return false
}
