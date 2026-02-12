package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aygumov-g/service-users-go/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	app, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal("err")
	}

	go app.Run()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	app.Shutdown(shutdownCtx)
}
