package log

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

type Metadatkey string

const (
	MetaNamespacedName Metadatkey = "namespaced-name"
)

func EnrichLoggerWithContext(ctx context.Context, logger logr.Logger) logr.Logger {
	newLogger := logger
	if ctx != nil {
		if namespacedName, ok := ctx.Value(MetaNamespacedName).(string); ok {
			newLogger = newLogger.WithValues("namespacedName", namespacedName)
		}
	}
	return newLogger
}

func NewDefaultZapLogger() (logr.Logger, error) {
	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		return logr.Discard(), err
	}
	logger := zapr.NewLogger(zapLogger)
	return logger, nil
}
