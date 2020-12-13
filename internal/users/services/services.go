package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/maxvoronov/tweetster/internal/users/models"
	"github.com/maxvoronov/tweetster/internal/users/repositories"
)

var ErrUserNotFound = status.Error(codes.NotFound, "User not found")

type UsersService interface {
	UserGetByID(ctx context.Context, id string) (*models.User, error)
}

type usersSvc struct {
	Storage repositories.UserRepository
}

func NewUsersService(storage repositories.UserRepository) UsersService {
	return &usersSvc{
		Storage: storage,
	}
}

func (svc usersSvc) UserGetByID(_ context.Context, id string) (*models.User, error) {
	user, err := svc.Storage.GetById(context.Background(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
