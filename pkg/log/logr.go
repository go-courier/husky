package log

import (
	"context"

	"github.com/go-logr/glogr"
	"github.com/go-logr/logr"
)

type contextKeyLogger int

var Logger = glogr.New()

func WithLogger(logger logr.Logger) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, contextKeyLogger(1), logger)
	}
}

func LoggerFromContext(ctx context.Context) logr.Logger {
	if logger, ok := ctx.Value(contextKeyLogger(1)).(logr.Logger); ok {
		return logger
	}
	return Logger
}
