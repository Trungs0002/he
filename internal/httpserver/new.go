package httpserver

import (
	pkgLog "gitlab.com/gma-vietnam/tanca-event/pkg/log"
	"gitlab.com/gma-vietnam/tanca-event/pkg/mongo"

	"github.com/gin-gonic/gin"
	pkgCrt "gitlab.com/gma-vietnam/tanca-event/pkg/encrypter"
)

type HTTPServer struct {
	gin          *gin.Engine
	l            pkgLog.Logger
	port         int
	database     mongo.Database
	jwtSecretKey string
	encrypter    pkgCrt.Encrypter
	secretConfig SecretConfig
}

type SecretConfig struct {
	SecretKey string
}

type Config struct {
	Port         int
	Database     mongo.Database
	JWTSecretKey string
	Encrypter    pkgCrt.Encrypter
	SecretConfig SecretConfig
}

func New(l pkgLog.Logger, cfg Config) *HTTPServer {
	return &HTTPServer{
		l:            l,
		gin:          gin.Default(),
		port:         cfg.Port,
		database:     cfg.Database,
		encrypter:    cfg.Encrypter,
		jwtSecretKey: cfg.JWTSecretKey,
		secretConfig: cfg.SecretConfig,
	}
}
