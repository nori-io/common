package transport

import (
	"context"

	"github.com/nori-io/nori-common/logger"
)

type ErrorHandler interface {
	Handle(ctx context.Context, err error)
}

type LogErrorHandler struct {
	logger logger.Logger
}

func NewLogErrorHandler(logger logger.Logger) *LogErrorHandler {
	return &LogErrorHandler{
		logger: logger,
	}
}

func (lh *LogErrorHandler) Handle(ctx context.Context, err error) {
	lh.logger.Error(err)
}
