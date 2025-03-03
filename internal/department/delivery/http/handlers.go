package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gma-vietnam/tanca-event/pkg/response"
)

func (h handler) create(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processCreateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "department.handler.create.processCreateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	b, err := h.uc.Create(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "department.handler.create.uc.Create: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(b))
}

func (h handler) listDepartment(c *gin.Context) {
	ctx := c.Request.Context()
	sc, err := h.processGetListRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "department.handler.listDepart.processGetRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}
	l, err := h.uc.GetAll(ctx, sc)
	if err != nil {
		h.l.Warnf(ctx, "department.handler.listDepart.uc.GetListDepart: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}
	response.OK(c, h.newListResp(l))
}
func (h handler) get(c *gin.Context) {
	ctx := c.Request.Context()

	sc, err := h.processGetRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "department.handler.get.processGetRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	b, err := h.uc.Get(ctx, sc)
	if err != nil {
		h.l.Warnf(ctx, "department.handler.get.uc.Get: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(b))
}

func (h handler) update(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processUpdateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "department.handler.update.processUpdateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	b, err := h.uc.Update(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "department.handler.update.uc.Update: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(b))
}

func (h handler) delete(c *gin.Context) {
	ctx := c.Request.Context()

	sc, err := h.processDeleteRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "department.handler.delete.processDeleteRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	err = h.uc.Delete(ctx, sc)
	if err != nil {
		h.l.Warnf(ctx, "department.handler.delete.uc.Delete: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, nil)
}