package services

import (
	"context"

	"github.com/maxvoronov/tweetster/internal/tweets/models"
)

type TweetsService interface {
	PostsGetList(ctx context.Context) ([]*models.Post, error)
	PostsGetById(ctx context.Context, id uint64) (*models.Post, error)
}
