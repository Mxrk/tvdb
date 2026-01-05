package tvdb

import (
	"errors"
	"fmt"
)

// APIError represents an error response from the TVDB API.
type APIError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("tvdb: API error %d %s: %s", e.StatusCode, e.Status, e.Body)
}

// IsNotFound returns true if the error is a 404 Not Found error.
func (e *APIError) IsNotFound() bool {
	return e.StatusCode == 404
}

// IsUnauthorized returns true if the error is a 401 Unauthorized error.
func (e *APIError) IsUnauthorized() bool {
	return e.StatusCode == 401
}

// IsBadRequest returns true if the error is a 400 Bad Request error.
func (e *APIError) IsBadRequest() bool {
	return e.StatusCode == 400
}

// IsAPIError checks if an error is an APIError and returns it.
func IsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}
