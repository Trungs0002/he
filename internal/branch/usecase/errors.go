package usecase

import (
	"errors"

	pkgErrors "gitlab.com/gma-vietnam/tanca-event/pkg/errors"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrRegionNotFound      = pkgErrors.NewHTTPError(404, "region not found")
	ErrBranchNotFound      = pkgErrors.NewHTTPError(404, "branch not found")
	ErrBranchHasDepartment = pkgErrors.NewHTTPError(403, "branch has department")
)

