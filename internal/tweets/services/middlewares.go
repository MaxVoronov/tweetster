package services

import (
	"context"

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

func (mw loggingMiddleware) PostsGetList(ctx context.Context) (_ []*models.Post, err error) {
	defer func() {
		_ = mw.logger.Log("svc", "", "method", "GetList", "err", err)
	}()
	return mw.next.PostsGetList(ctx)
}

func (mw loggingMiddleware) PostsGetById(ctx context.Context, id uint64) (_ *models.Post, err error) {
	defer func() {
		_ = mw.logger.Log("svc", "", "method", "GetById", "id", id, "err", err)
	}()
	return mw.next.PostsGetById(ctx, id)
}
