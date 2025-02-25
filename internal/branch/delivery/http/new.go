package http

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gma-vietnam/tanca-event/internal/branch"
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

type Handler interface {
	create(c *gin.Context)
}

type handler struct {
	l  log.Logger
	uc branch.Usecase
}

// New returns a new instance of the HTTPHandler interface
func New(l log.Logger, uc branch.Usecase) Handler {
	return handler{
		l:  l,
		uc: uc,
	}
}
