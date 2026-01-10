package app

import (
	"context"

	"github.com/aygumov-g/service-users-go/internal/config"
	"github.com/aygumov-g/service-users-go/internal/http/server"
	"github.com/aygumov-g/service-users-go/internal/logger"
	"github.com/aygumov-g/service-users-go/internal/logger/slog"
)

type App struct {
	httpServer *server.Server
	logger     logger.Logger
}

func New(ctx context.Context) (*App, error) {
	cfg := config.Load()
	log := slog.New()

	httpServer := buildHTTP(cfg)

	return &App{
		httpServer: httpServer,
		logger:     log,
	}, nil
}
