package apperrs

import (
	"net/http"

	"github.com/teambition/gear"
)

// Predefined errors.
var (
	// ErrForbidden ...
	ErrForbidden = &gear.Error{Code: http.StatusForbidden, Err: "Forbidden"}
)
