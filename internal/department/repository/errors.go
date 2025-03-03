package repository

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrBranchNotFound     = errors.New("branch not found")
	ErrDepartmentNotFound = errors.New("department not found")
)