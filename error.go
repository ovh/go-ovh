package govh

import "fmt"

// APIOvhError represents an error that can occured while calling the API.
type APIOvhError struct {
	// Error message.
	Message string
	// HTTP code.
	Code int
}

func (err *APIOvhError) Error() string {
	return fmt.Sprintf("Error %d : %q", err.Code, err.Message)
}
