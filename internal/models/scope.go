package models

type Scope struct {
	UserID       string `json:"user_id"`
	Role         string `json:"role"`
	ShopID       string `json:"shop_id"`
	RegionID     string `json:"region_id"`
	BranchID     string `json:"branch_id"`
	DepartmentID string `json:"department_id"`
}
