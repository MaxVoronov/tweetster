package services

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/maxvoronov/tweetster/internal/users/models"
)

// TODO: Move to special directory?
func LoggingMiddleware(logger log.Logger) func(service UsersService) UsersService {
	return func(next UsersService) UsersService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   UsersService
}

func (mw loggingMiddleware) UserGetByID(ctx context.Context, id uint64) (*models.User, error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log("svc", "users", "method", "UserGetByID", "id", id, "time", time.Since(begin).Seconds())
	}(time.Now())
	return mw.next.UserGetByID(ctx, id)
}
