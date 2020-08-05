package gobikeindex

// AccessToken holds everything required to aOAth
type AccessToken struct {
	ID          string
	Secret      string
	Name        string
	CallbackURL *string
}
