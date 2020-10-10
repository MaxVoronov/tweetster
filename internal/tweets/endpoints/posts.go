package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/maxvoronov/tweetster/internal/tweets/models"
	"github.com/maxvoronov/tweetster/internal/tweets/services"
)

// Posts: Get List
type PostsGetListRequest struct{}

type PostsGetListResponse struct {
	Posts []*models.Post
}

func makePostsGetListEndpoint(svc services.TweetsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		posts, err := svc.PostsGetList(ctx)
		return &PostsGetListResponse{Posts: posts}, err
	}
}

// Posts: Get by ID
type PostsGetByIDRequest struct {
	ID uint64
}

type PostsGetByIDResponse struct {
	Post *models.Post
}

func makePostsGetByIDEndpoint(svc services.TweetsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(*PostsGetByIDRequest).ID
		post, err := svc.PostsGetByID(ctx, id)
		if err != nil {
			return nil, err
		}

		return &PostsGetByIDResponse{Post: post}, nil
	}
}
