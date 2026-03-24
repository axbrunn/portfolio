package blog

import "errors"

var ErrNoRecord = errors.New("no matching record found")

// ValidationError wordt teruggegeven als validatie faalt.
// Bevat fouten per veld zodat de handler ze in de form kan tonen.
type ValidationError struct {
	Fields map[string]string
}

func (e *ValidationError) Error() string {
	return "validation failed"
}
