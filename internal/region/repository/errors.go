package repository

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrShopNotFound   = errors.New("shop not found")
	ErrRegionNotFound = errors.New("region not found")
)