package endpoints

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/maxvoronov/tweetster/internal/users/services"
)

type Endpoints struct {
	UsersGetByIDEndpoint endpoint.Endpoint
}

func PrepareServiceEndpoints(svc services.UsersService) Endpoints {
	return Endpoints{
		UsersGetByIDEndpoint: makeUsersGetByIDEndpoint(svc),
	}
}
