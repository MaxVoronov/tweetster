package middlewares

import (
	"context"
	"time"

	"github.com/go-logr/logr"

	"github.com/maxvoronov/tweetster/internal/users/models"
	"github.com/maxvoronov/tweetster/internal/users/services"
)

func LoggingMiddleware(logger logr.Logger) func(service services.UsersService) services.UsersService {
	return func(next services.UsersService) services.UsersService {
		return loggingMiddleware{next, logger}
	}
}

type loggingMiddleware struct {
	next   services.UsersService
	logger logr.Logger
}

func (mw loggingMiddleware) UserGetByID(ctx context.Context, id uint64) (*models.User, error) {
	defer func(begin time.Time) {
		mw.logger.Info("Call users.UserGetByID", "id", id, "latency", time.Since(begin).Nanoseconds())
	}(time.Now())

	return mw.next.UserGetByID(ctx, id)
}
