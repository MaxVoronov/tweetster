package repositories

import (
	"context"

	"github.com/maxvoronov/tweetster/internal/users/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetById(ctx context.Context, id string) (*models.User, error)
}
