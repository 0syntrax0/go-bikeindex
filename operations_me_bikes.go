package gobikeindex

// OperationsMeBike Current user's bikes
type OperationsMeBike struct {
}

// OperationsMeBikes list
type OperationsMeBikes []OperationsMeBike

// OperationsMeBikes - This returns the current user’s bikes, so long as the access_token has the `read_bikes` scope.
// This uses the bike list bike objects, which only contains the most important information.
// To get all possible information about a bike use `/bikes/{id}`
func (bi *BikeIndex) OperationsMeBikes() (OperationsMeBikes, error) {
	opBikes := OperationsMeBikes{}
	err := bi.getJSON(bi.buildURL("me", "bikes"), &opBikes)
	return opBikes, err
}
