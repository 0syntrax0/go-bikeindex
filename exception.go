package gobikeindex

import "fmt"

// ExceptionCode -
type ExceptionCode int

// Exception is what's returned when an error is encountered
type Exception struct {
	Status  int    // HTTP error code
	Message string // Returned error message
	// Error   string // error message from bikeindex
}

func (e Exception) Error() string {
	var status int
	if e.Status != status {
		return fmt.Sprintf("Status %d: %s", e.Status, e.Message)
	}
	return e.Message
}
