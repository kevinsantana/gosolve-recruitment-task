package share

import (
	"errors"
	"fmt"
	"strings"
)

// internal infra errors
var (
	// rest
	ErrTimeout     = errors.New("error timeout")
	ErrContentType = errors.New("content-type error")
	ErrValidation  = errors.New("validation error")
)

type DomainError struct {
	Domain      string
	Module      string
	Err         string
	Description string
}

func (err DomainError) Error() string {
	return fmt.Sprintf(
		"%s|%s|%s",
		strings.ToUpper(err.Domain),
		strings.ToUpper(err.Module),
		strings.ToUpper(err.Err),
	)
}

type ClientError struct {
	Domain      string
	Module      string
	Err         string
	Description string
}

func (err ClientError) Error() string {
	return fmt.Sprintf(
		"%s|%s|%s",
		strings.ToUpper(err.Domain),
		strings.ToUpper(err.Module),
		strings.ToUpper(err.Err),
	)
}
