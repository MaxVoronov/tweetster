package middlewares

import (
	"context"
	"time"

	"github.com/go-logr/logr"

	"github.com/maxvoronov/tweetster/internal/tweets/models"
	"github.com/maxvoronov/tweetster/internal/tweets/services"
)

func LoggingMiddleware(logger logr.Logger) func(service services.TweetsService) services.TweetsService {
	return func(next services.TweetsService) services.TweetsService {
		return loggingMiddleware{next, logger}
	}
}

type loggingMiddleware struct {
	next   services.TweetsService
	logger logr.Logger
}

func (mw loggingMiddleware) PostsGetList(ctx context.Context) ([]*models.Post, error) {
	defer func(begin time.Time) {
		mw.logger.Info("Call tweets.PostsGetList", "latency", time.Since(begin).Nanoseconds())
	}(time.Now())
	return mw.next.PostsGetList(ctx)
}

func (mw loggingMiddleware) PostsGetByID(ctx context.Context, id uint64) (*models.Post, error) {
	defer func(begin time.Time) {
		mw.logger.Info("Call tweets.PostsGetByID", "id", id, "latency", time.Since(begin).Nanoseconds())
	}(time.Now())
	return mw.next.PostsGetByID(ctx, id)
}
