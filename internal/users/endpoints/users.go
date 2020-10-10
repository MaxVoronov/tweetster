package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/maxvoronov/tweetster/internal/users/models"
	"github.com/maxvoronov/tweetster/internal/users/services"
)

// Users: Get by ID
type UsersGetByIDRequest struct {
	ID uint64
}

type UsersGetByIDResponse struct {
	User *models.User
}

func makeUsersGetByIDEndpoint(svc services.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(*UsersGetByIDRequest).ID
		user, err := svc.UserGetByID(ctx, id)
		if err != nil {
			return nil, err
		}

		return &UsersGetByIDResponse{User: user}, nil
	}
}
