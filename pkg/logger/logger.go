package logger

import (
	"emvn/config"
	"io"
	"log/slog"
	"os"
)

var LogLelvel = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

// Because time limit, I will not implement any custom logger with request id or trace,....
// TODO: Custom logger with ctx and trace
func NewLogger(logConfig config.LogConfig) {
	opts := slog.HandlerOptions{
		Level:     LogLelvel[logConfig.Level],
		AddSource: true,
	}

	// TODO: Write to file, log rotation
	writer := io.MultiWriter(os.Stdout)

	handler := slog.NewTextHandler(writer, &opts)

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
