package app

import (
	"net/http"
	"time"

	"github.com/aygumov-g/service-users-go/internal/config"
	"github.com/aygumov-g/service-users-go/internal/integration/sso"
	user_repo "github.com/aygumov-g/service-users-go/internal/repository/postgres/user"
	user_srv "github.com/aygumov-g/service-users-go/internal/service/user"
	me_handler "github.com/aygumov-g/service-users-go/internal/transport/http/handler/me"
	users_handler "github.com/aygumov-g/service-users-go/internal/transport/http/handler/users"
	auth_id "github.com/aygumov-g/service-users-go/internal/transport/http/identity/auth"
	auth_mw "github.com/aygumov-g/service-users-go/internal/transport/http/middleware/auth"
	methods_mw "github.com/aygumov-g/service-users-go/internal/transport/http/middleware/method"
	"github.com/aygumov-g/service-users-go/internal/transport/http/router"
	"github.com/aygumov-g/service-users-go/internal/transport/http/server"
	"github.com/aygumov-g/service-users-go/pkg/clock"
	"github.com/aygumov-g/service-users-go/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

func buildHTTP(cfg *config.Config, db *pgxpool.Pool, log logger.Logger) *server.Server {
	ssoClient := sso.NewSSO(
		"http://service-sso-app-go:39179/auth/me",
		5*time.Second,
	)

	clk := clock.NewSystemClock()

	userRepo := user_repo.NewRepository(db)
	userService := user_srv.NewService(userRepo, clk)

	authIdentity := auth_id.NewIdentity("identity")

	authMW := auth_mw.NewMiddleware(ssoClient, authIdentity)
	methodsMW := methods_mw.NewMiddleware()

	meHandler := me_handler.NewHandler(userService, authIdentity)
	usersHandler := users_handler.NewHandler(userService, authIdentity)

	r := router.NewRouter()
	r.Handle("/users", methodsMW.Handle([]string{http.MethodGet}, authMW.Handle(usersHandler)))
	r.Handle("/users/me", methodsMW.Handle([]string{http.MethodGet, http.MethodPatch}, authMW.Handle(meHandler)))

	return server.NewServer(cfg.AppPort, r.Handler())
}
