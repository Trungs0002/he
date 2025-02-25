package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gma-vietnam/tanca-event/pkg/response"
)

func (h handler) create(c *gin.Context) {
	ctx := c.Request.Context()

	req, sc, err := h.processCreateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.create.processCreateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	b, err := h.uc.Create(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.create.uc.Create: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(b))
}
