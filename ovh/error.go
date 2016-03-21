package ovh

import "fmt"

// APIError represents an error that can occured while calling the API.
type APIError struct {
	// Error message.
	Message string
	// HTTP code.
	Code int
}

func (err *APIError) Error() string {
	return fmt.Sprintf("Error %d: %q", err.Code, err.Message)
}
