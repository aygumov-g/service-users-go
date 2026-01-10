package slog

import (
	"log/slog"
	"os"

	"github.com/aygumov-g/service-users-go/internal/logger"
)

type Logger struct {
	log *slog.Logger
}

func New() logger.Logger {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	return &Logger{log: slog.New(handler)}
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}
