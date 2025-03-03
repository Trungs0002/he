package http

import (
	"strings"

	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	"gitlab.com/gma-vietnam/tanca-event/pkg/response"
)

type createReq struct {
	Name  string `json:"name" binding:"required"`
	Alias string `json:"alias" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func (r createReq) validate() error {
	if strings.TrimSpace(r.Name) == "" || strings.TrimSpace(r.Alias) == "" || strings.TrimSpace(r.Code) == "" {
		return errWrongBody
	}
	return nil
}

func (r createReq) toInput() branch.CreateInput {
	return branch.CreateInput{
		Name:  r.Name,
		Alias: r.Alias,
		Code:  r.Code,
	}
}

type getQuery struct {
	Id string `form:"id"`
}

func (r getQuery) validate() error {
	if strings.TrimSpace(r.Id) == "" {
		return errWrongQuery
	}
	return nil
}

type updateReq struct {
	Name  string `json:"name"`
	Alias string `json:"alias"`
	Code  string `json:"code"`
}

func (r updateReq) validate() error {
	if strings.TrimSpace(r.Name) == "" && strings.TrimSpace(r.Alias) == "" && strings.TrimSpace(r.Code) == "" {
		return errWrongBody
	}
	return nil
}

func (r updateReq) toInput() branch.UpdateInput {
	return branch.UpdateInput{
		Name:  r.Name,
		Alias: r.Alias,
		Code:  r.Code,
	}
}

type query struct {
	ID string `json:"id" form:"id"`
}

type detailResp struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Alias     string            `json:"alias"`
	Code      string            `json:"code"`
	CreatedAt response.DateTime `json:"created_at"`
}
func (h handler) newDetailResp(d models.Branch) detailResp {
	return detailResp{
		ID:        d.ID.Hex(),
		Name:      d.Name,
		Alias:     d.Alias,
		Code:      d.Code,
		CreatedAt: response.DateTime(d.CreatedAt),
	}
}

type listResp struct {
	Branches []detailResp `json:"branches"`
}

func (h handler) newListResp(d []models.Branch) listResp {
	var branches []detailResp
	for _, branch := range d {
		branches = append(branches, h.newDetailResp(branch))
	}
	return listResp{
		Branches: branches,
	}
}