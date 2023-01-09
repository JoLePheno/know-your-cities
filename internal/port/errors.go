package port

import (
	"errors"
	"io"
)

var (
	ErrEOF               = io.EOF
	ErrInvalidData       = errors.New("invalid set of data")
	ErrInvalidZip        = errors.New("zip code not valid")
	ErrZipCodeNotFound   = errors.New("zip code not found")
	ErrInternalError     = errors.New("internal error")
	ErrPgCityNotFound    = errors.New("postgres error city not found")
	ErrCityAlreadyStored = errors.New("city already stored in database")
)
