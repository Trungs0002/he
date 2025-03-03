package http

import (
	pkgErrors "gitlab.com/gma-vietnam/tanca-event/pkg/errors"
)

var (
	errWrongBody  = pkgErrors.NewHTTPError(10000, "Wrong body")
	errWrongQuery = pkgErrors.NewHTTPError(10001, "Wrong query")
)

func (h handler) mapError(err error) error {
	return err
}

