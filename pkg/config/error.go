package config

import "errors"

var (
	// ErrNotFound is thrown when the config is not found
	ErrNotFound = errors.New("config not found")
)
