package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gma-vietnam/tanca-event/internal/models"
	pkgErrors "gitlab.com/gma-vietnam/tanca-event/pkg/errors"
	"gitlab.com/gma-vietnam/tanca-event/pkg/jwt"
)

func (h handler) processCreateRequest(c *gin.Context) (createReq, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "event.region.http.processCreateRequest.GetPayloadFromContext: unauthorized")
		return createReq{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "event.region.http.processCreateRequest.ShouldBindJSON: %v", err)
		return createReq{}, models.Scope{}, errWrongBody
	}

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "event.region.http.processCreateRequest.validate: %v", err)
		return createReq{}, models.Scope{}, err
	}

	sc := jwt.NewScope(payload)

	return req, sc, nil
}

func (h handler) processGetListRequest(c *gin.Context) (models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "event.region.http.processGetRequest.GetPayloadFromContext: unauthorized")
		return models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}
	var query getQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		h.l.Warnf(ctx, "event.region.http.processGetRequest.ShouldBindQuery: %v", err)
		return models.Scope{}, errWrongQuery
	}
	sc := jwt.NewScope(payload)

	return sc, nil
}
func (h handler) processGetRequest(c *gin.Context) (models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "event.region.http.processGetRequest.GetPayloadFromContext: unauthorized")
		return models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	id := c.Param("id")
	sc := jwt.NewScope(payload)
	sc.RegionID = id
	return sc, nil
}
func (h handler) processUpdateRequest(c *gin.Context) (updateReq, models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "event.region.http.processUpdateRequest.GetPayloadFromContext: unauthorized")
		return updateReq{}, models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	var req updateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Warnf(ctx, "event.region.http.processUpdateRequest.ShouldBindJSON: %v", err)
		return updateReq{}, models.Scope{}, errWrongBody
	}

	id := c.Param("id")

	if err := req.validate(); err != nil {
		h.l.Warnf(ctx, "event.region.http.processUpdateRequest.validate: %v", err)
		return updateReq{}, models.Scope{}, err
	}

	sc := jwt.NewScope(payload)
	sc.RegionID = id
	return req, sc, nil
}

func (h handler) processDeleteRequest(c *gin.Context) (models.Scope, error) {
	ctx := c.Request.Context()

	payload, ok := jwt.GetPayloadFromContext(ctx)
	if !ok {
		h.l.Warnf(ctx, "event.region.http.processDeleteRequest.GetPayloadFromContext: unauthorized")
		return models.Scope{}, pkgErrors.NewUnauthorizedHTTPError()
	}

	id := c.Param("id")

	sc := jwt.NewScope(payload)
	sc.RegionID = id

	return sc, nil
}