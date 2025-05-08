package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func Init() {
	opts := slog.HandlerOptions{AddSource: false, Level: slog.LevelInfo}
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &opts))
}

func Logger() *slog.Logger {
	return logger
}
