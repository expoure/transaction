package customized_errors

import "errors"

var (
	EntityNotFound = errors.New("Entity not found")
	DuplicateKey   = errors.New("Duplicate key")
)
