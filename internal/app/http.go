package app

import (
	"github.com/aygumov-g/service-users-go/internal/config"
	"github.com/aygumov-g/service-users-go/internal/http/router"
	"github.com/aygumov-g/service-users-go/internal/http/server"
	userservice "github.com/aygumov-g/service-users-go/internal/service/user"
	userstorage "github.com/aygumov-g/service-users-go/internal/storage/postgres/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

func buildHTTP(cfg *config.Config, pool *pgxpool.Pool) *server.Server {
	userRepo := userstorage.New(pool)
	_ = userservice.New(userRepo)

	r := router.New()

	return server.New(":"+cfg.AppPort, r.Handler())
}
