package middleware

import (
	"gitlab.com/gma-vietnam/tanca-event/pkg/log"

	pkgCrt "gitlab.com/gma-vietnam/tanca-event/pkg/encrypter"
	"gitlab.com/gma-vietnam/tanca-event/pkg/jwt"
)

type Middleware struct {
	l          log.Logger
	jwtManager jwt.Manager
	encrypter  pkgCrt.Encrypter
}

func New(l log.Logger, jwtManager jwt.Manager, encrypter pkgCrt.Encrypter) Middleware {
	return Middleware{
		l:          l,
		jwtManager: jwtManager,
		encrypter:  encrypter,
	}
}
