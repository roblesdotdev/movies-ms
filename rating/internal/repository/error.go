package repository

import "errors"

var (
	ErrEntryNotFound = errors.New("entry not found")
	ErrNotFound      = errors.New("not found")
)
