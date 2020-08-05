package gobikeindex

// PublicImage public image ¯\_(ツ)_/¯
type PublicImage struct {
	Name   string `json:"name"`
	Full   string `json:"full"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Thumb  string `json:"thumb"`
	ID     int    `json:"id"`
}

// PublicImages -
type PublicImages []*PublicImage
