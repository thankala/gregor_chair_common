package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

// Get returns the shared logger instance
func Get() *slog.Logger {
	return logger
}
