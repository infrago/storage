package storage

import "errors"

const (
	NAME = "STORE"
)

var (
	errInvalidConnection = errors.New("Invalid storage connection.")
	errInvalidCode       = errors.New("Invalid code.")
	errInvalidHandler    = errors.New("Invalid storage handler.")
)
