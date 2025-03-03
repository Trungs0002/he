package usecase

import (
	"errors"

	pkgErrors "gitlab.com/gma-vietnam/tanca-event/pkg/errors"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrBranchNotFound      = pkgErrors.NewHTTPError(404, "branch not found")
	ErrDepartmentNotFound  = pkgErrors.NewHTTPError(404, "department not found")
	ErrDepartmentHasBranch = pkgErrors.NewHTTPError(403, "department has branch")
)