package app

import (
	"github.com/aygumov-g/service-users-go/internal/config"
	"github.com/aygumov-g/service-users-go/internal/http/router"
	"github.com/aygumov-g/service-users-go/internal/http/server"
)

func buildHTTP(cfg *config.Config) *server.Server {
	r := router.New()

	return server.New(":"+cfg.AppPort, r.Handler())
}
