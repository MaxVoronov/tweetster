package services

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/maxvoronov/tweetster/internal/users/models"
)

var ErrUserNotFound = status.Error(codes.NotFound, "User not found")

type usersSvc struct {
	Storage []*models.User
}

func NewUsersService() UsersService {
	users := make([]*models.User, 0, 1)
	users = append(users, &models.User{
		ID:    1,
		Login: "tester",
		Email: "tester@email.com",
		Name:  "Just Tester",
	})

	return &usersSvc{Storage: users}
}

func (svc usersSvc) UserGetByID(_ context.Context, id uint64) (*models.User, error) {
	for _, user := range svc.Storage {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, ErrUserNotFound
}
