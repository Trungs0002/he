package repository

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrRegionNotFound = errors.New("region not found")
	ErrBranchNotFound = errors.New("branch not found")
)
