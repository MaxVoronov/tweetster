package services

import (
	"context"

	"github.com/maxvoronov/tweetster/internal/users/models"
)

type UsersService interface {
	UserGetByID(ctx context.Context, id uint64) (*models.User, error)
}
