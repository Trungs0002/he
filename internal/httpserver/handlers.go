package httpserver

import (
	branchHTTP "gitlab.com/gma-vietnam/tanca-event/internal/branch/delivery/http"
	branchMongo "gitlab.com/gma-vietnam/tanca-event/internal/branch/repository/mongo"
	branchUsecase "gitlab.com/gma-vietnam/tanca-event/internal/branch/usecase"
)

func (srv HTTPServer) mapHandlers() {
	// Repositories
	branchRepo := branchMongo.NewRepository(srv.l, srv.database)

	// Usecases
	branchUC := branchUsecase.New(srv.l, branchRepo)

	// Handlers
	branchH := branchHTTP.New(srv.l, branchUC)

	// Middlewares

	api := srv.gin.Group("/api/v1")

	branchHTTP.MapRoutes(api.Group("/branches"), mw, branchH)
}
