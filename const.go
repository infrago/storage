package store

import "errors"

const (
	NAME = "STORE"
)

var (
	errInvalidStoreConnection = errors.New("Invalid store connection.")
	errInvalidCode            = errors.New("Invalid code.")
	errInvalidHandler         = errors.New("Invalid store handler.")
)
