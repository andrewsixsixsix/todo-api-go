package logger

import (
	"log/slog"
	"os"
)

// TODO: make private and add getter func
var Logger *slog.Logger

func Init() *slog.Logger {
	opts := slog.HandlerOptions{AddSource: false, Level: slog.LevelInfo}
	return slog.New(slog.NewJSONHandler(os.Stdout, &opts))
}
