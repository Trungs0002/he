package http

import (
	pkgErrors "gitlab.com/gma-vietnam/tanca-event/pkg/errors"
)

var (
	errWrongBody = pkgErrors.NewHTTPError(10000, "Wrong body")
)

func (h handler) mapError(err error) error {
	return err
}
