package http

import (
	"github.com/gin-gonic/gin"
	middleware "gitlab.com/gma-vietnam/tanca-event/internal/middleware"
)

// MapRoutes maps the routes to the handler functions
func MapRoutes(r *gin.RouterGroup, mw middleware.Middleware, h Handler) {
	r.Use(mw.Auth())
	r.POST("", h.create)
	r.GET("", h.listDepartment)
	r.GET("/:id", h.get)
	r.PATCH("/:id", h.update)
	r.DELETE("/:id", h.delete)
}