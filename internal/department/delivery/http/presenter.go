package http

import (
	"strings"

	"gitlab.com/gma-vietnam/tanca-event/internal/department"
	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	"gitlab.com/gma-vietnam/tanca-event/pkg/response"
)

type createReq struct {
	Name string `json:"name" binding:"required"`
}

func (r createReq) validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	return nil
}

func (r createReq) toInput() department.CreateInput {
	return department.CreateInput{
		Name: r.Name,
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
	Name string `json:"name"`
}

func (r updateReq) validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	return nil
}

func (r updateReq) toInput() department.UpdateInput {
	return department.UpdateInput{
		Name: r.Name,
	}
}

type query struct {
	ID string `json:"id" form:"id"`
}

type detailResp struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	CreatedAt response.DateTime `json:"created_at"`
}

func (h handler) newDetailResp(d models.Department) detailResp {
	return detailResp{
		ID:        d.ID.Hex(),
		Name:      d.Name,
		CreatedAt: response.DateTime(d.CreatedAt),
	}
}

type listResp struct {
	Departments []detailResp `json:"departmentes"`
}

func (h handler) newListResp(d []models.Department) listResp {
	var departmentes []detailResp
	for _, department := range d {
		departmentes = append(departmentes, h.newDetailResp(department))
	}
	return listResp{
		Departments: departmentes,
	}
}