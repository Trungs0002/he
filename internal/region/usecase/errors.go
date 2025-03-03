package usecase

import (
	"errors"

	pkgErrors "gitlab.com/gma-vietnam/tanca-event/pkg/errors"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrShopNotFound    = pkgErrors.NewHTTPError(404, "shop not found")
	ErrRegionNotFound  = pkgErrors.NewHTTPError(404, "region not found")
	ErrRegionHasBranch = pkgErrors.NewHTTPError(403, "region has branch")
)