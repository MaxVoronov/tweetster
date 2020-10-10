package services

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/maxvoronov/tweetster/internal/tweets/models"
)

// TODO: Move to special directory?
func LoggingMiddleware(logger log.Logger) func(TweetsService) TweetsService {
	return func(next TweetsService) TweetsService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   TweetsService
}

func (mw loggingMiddleware) PostsGetList(ctx context.Context) ([]*models.Post, error) {
	defer func(begin time.Time) {
		mw.logger.Log("svc", "tweets", "method", "PostsGetList", "time", time.Since(begin).Seconds())
	}(time.Now())
	return mw.next.PostsGetList(ctx)
}

func (mw loggingMiddleware) PostsGetByID(ctx context.Context, id uint64) (*models.Post, error) {
	defer func(begin time.Time) {
		mw.logger.Log("svc", "tweets", "method", "PostsGetByID", "id", id, "time", time.Since(begin).Seconds())
	}(time.Now())
	return mw.next.PostsGetByID(ctx, id)
}
