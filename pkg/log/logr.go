package log

import (
	"context"
	stdlog "log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

type contextKeyLogger int

func SetVerbosity(v int) {
	stdr.SetVerbosity(v)
}

var Logger = stdr.New(stdlog.New(os.Stderr, "[husky] ", stdlog.Ltime))

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
