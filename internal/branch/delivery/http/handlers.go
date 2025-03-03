package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gma-vietnam/tanca-event/pkg/response"
)

func (h handler) create(c *gin.Context) { //tao branch moi
	ctx := c.Request.Context()
	req, sc, err := h.processCreateRequest(c) //goi ham processCreateRequest
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.create.processCreateRequest: %s", err)
		mapErr := h.mapError(err) //tra loi neu co loi
		response.Error(c, mapErr)
		return
	}
	b, err := h.uc.Create(ctx, sc, req.toInput()) //tao moi
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.create.uc.Create: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(b)) //tra ket qua
}

func (h handler) listBranches(c *gin.Context) { //lay danh sach branch
	ctx := c.Request.Context()
	sc, err := h.processGetListRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.listDepart.processGetRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}
	l, err := h.uc.GetAll(ctx, sc)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.listDepart.uc.GetListDepart: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}
	response.OK(c, h.newListResp(l))
}
func (h handler) get(c *gin.Context) { //lay thong tin branch
	ctx := c.Request.Context()

	sc, err := h.processGetRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.get.processGetRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	b, err := h.uc.Get(ctx, sc)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.get.uc.Get: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(b))
}

func (h handler) update(c *gin.Context) { //cap nhat thong tin branch
	ctx := c.Request.Context()

	req, sc, err := h.processUpdateRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.update.processUpdateRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	b, err := h.uc.Update(ctx, sc, req.toInput())
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.update.uc.Update: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, h.newDetailResp(b))
}

func (h handler) delete(c *gin.Context) { //xoa branch
	ctx := c.Request.Context()

	sc, err := h.processDeleteRequest(c)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.delete.processDeleteRequest: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	err = h.uc.Delete(ctx, sc)
	if err != nil {
		h.l.Warnf(ctx, "branch.handler.delete.uc.Delete: %s", err)
		mapErr := h.mapError(err)
		response.Error(c, mapErr)
		return
	}

	response.OK(c, nil)
}

// func (srv HTTPServer) mapHandlers() {
// 	jwtManager := jwt.NewManager(srv.jwtSecretKey)
// 	// Repositories
// 	branchRepo := branchMongo.NewRepository(srv.l, srv.database)

// 	userRepo := userMongo.NewRepository(srv.l, srv.database)
// 	// Usecases
// 	branchUC := branchUsecase.New(srv.l, branchRepo)

// 	userUC := userUsecase.New(srv.l, userRepo, jwtManager)
// 	// Handlers
// 	branchH := branchHTTP.New(srv.l, branchUC)
// 	userH := userHTTP.New(srv.l, userUC)

// 	// Middlewares
// 	mw := middleware.New(srv.l, jwtManager, srv.encrypter)
// 	api := srv.gin.Group("/api/v1")

// 	userHTTP.MapRoutes(api.Group("/users"), userH)
// 	branchHTTP.MapRoutes(api.Group("/branches"), mw, branchH)
// }