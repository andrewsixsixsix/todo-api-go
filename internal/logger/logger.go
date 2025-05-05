package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func Init() *slog.Logger {
	opts := slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}
	return slog.New(slog.NewJSONHandler(os.Stdout, &opts))
}
