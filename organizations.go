package gobikeindex

// OrganizationReq -
type OrganizationReq struct {
	// required fields
	Name           string
	Website        string
	LocationName   []string
	LocationStreet []string
	LocationCity   []string

	// optional
	Kind            *string
	LocationState   *[]string
	LocationCountry *[]string
	LocationZip     *[]string
	LocationPhone   *[]string
}

// Organization - Note: Access to this endpoint is only available to select api clients.
type Organization struct {
}
