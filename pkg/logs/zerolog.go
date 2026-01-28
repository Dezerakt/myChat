package logs

import (
	"context"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func FromContext(ctx context.Context) *Logger {
	logger := zerolog.Ctx(ctx)
	return &Logger{logger}
}
