package gobikeindex

// OperationsMe - Current user is the owner of the `access_token` you use in the request. Depending on your scopes you will get different things back.
// You will always get the user’s `id`
// For an array of the user’s bike ids, you need `read_bikes` access
// For a hash of information about the user (including their email address), you need `read_user` access
// For an array of the organizations and/or shops they’re a part of, `read_organization_membership` access
// TODO: fix after OAuth2 implementation
type OperationsMe struct {
	ID int64 `json:"id"`
}

// OperationsMe - Current user is the owner of the `access_token` you use in the request. Depending on your scopes you will get different things back.
func (bi *BikeIndex) OperationsMe() (OperationsMe, error) {
	opMe := OperationsMe{}
	err := bi.getJSON(bi.buildURL("me"), &opMe)
	return opMe, err
}
