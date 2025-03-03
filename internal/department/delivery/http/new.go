package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gma-vietnam/tanca-event/internal/department"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

type Handler interface {
	create(c *gin.Context)
	listDepartment(c *gin.Context)
	update(c *gin.Context)
	delete(c *gin.Context)
	get(c *gin.Context)
}

type handler struct {
	l  log.Logger
	uc department.Usecase
}

func New(l log.Logger, uc department.Usecase) Handler {
	return handler{
		l:  l,
		uc: uc,
	}
}