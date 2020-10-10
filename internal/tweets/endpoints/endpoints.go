package endpoints

import (
	"github.com/go-kit/kit/endpoint"

	"github.com/maxvoronov/tweetster/internal/tweets/services"
)

type Endpoints struct {
	PostsGetListEndpoint endpoint.Endpoint
	PostsGetByIDEndpoint endpoint.Endpoint
}

func PrepareServiceEndpoints(svc services.TweetsService) Endpoints {
	return Endpoints{
		PostsGetListEndpoint: makePostsGetListEndpoint(svc),
		PostsGetByIDEndpoint: makePostsGetByIDEndpoint(svc),
	}
}
