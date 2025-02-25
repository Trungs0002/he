package main

import (
	"gitlab.com/gma-vietnam/tanca-event/config"
	"gitlab.com/gma-vietnam/tanca-event/internal/appconfig/mongo"
	"gitlab.com/gma-vietnam/tanca-event/internal/httpserver"
	pkgLog "gitlab.com/gma-vietnam/tanca-event/pkg/log"
)

// @title Tanca Intern Golang API
// @description This is the API documentation for Tanca Intern Golang.

// @version 1
// @host api.tanca.io/
// @schemes https

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	client, err := mongo.Connect(cfg.Mongo.URI)
	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect(client)

	db := client.Database(cfg.Mongo.DBName)

	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	srv := httpserver.New(l, httpserver.Config{
		Port:         cfg.HTTPServer.Port,
		Database:     db,
		JWTSecretKey: cfg.JWT.SecretKey,
	})
	srv.Run()
}
