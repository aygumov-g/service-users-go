package app

import "context"

func (a *App) Shutdown(ctx context.Context) {
	a.logger.Info("shutdown started")

	_ = a.httpServer.Shutdown(ctx)

	a.logger.Info("shutdown completed")
}
