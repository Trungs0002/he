package http

import (
	"strings"

	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	"gitlab.com/gma-vietnam/tanca-event/internal/region"
	"gitlab.com/gma-vietnam/tanca-event/pkg/response"
)

// CreateReq is the request body for creating a new region
type createReq struct {
	Name string `json:"name" binding:"required"`
}

func (r createReq) validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	return nil
}

func (r createReq) toInput() region.CreateInput {
	return region.CreateInput{
		Name: r.Name,
	}
}

// GetQuery is the query string
type getQuery struct {
	Id string `form:"id"`
}

func (r getQuery) validate() error {
	if strings.TrimSpace(r.Id) == "" {
		return errWrongQuery
	}
	return nil
}

// UpdateReq is the request body for updating a region
type updateReq struct {
	Name string `json:"name"`
}

func (r updateReq) validate() error {
	if strings.TrimSpace(r.Name) == "" {
		return errWrongBody
	}
	return nil
}

func (r updateReq) toInput() region.UpdateInput {
	return region.UpdateInput{
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

func (h handler) newDetailResp(d models.Region) detailResp {
	return detailResp{
		ID:        d.ID.Hex(),
		Name:      d.Name,
		CreatedAt: response.DateTime(d.CreatedAt),
	}
}

type listResp struct {
	Regions []detailResp `json:"regiones"`
}

func (h handler) newListResp(d []models.Region) listResp {
	var regiones []detailResp
	for _, region := range d {
		regiones = append(regiones, h.newDetailResp(region))
	}
	return listResp{
		Regions: regiones,
	}
}