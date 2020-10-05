package services

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/maxvoronov/tweetster/internal/tweets/models"
)

var ErrPostNotFound = status.Error(codes.NotFound, "Post not found")

type tweetsSvc struct {
	Storage []*models.Post
}

func NewTweetsService() TweetsService {
	posts := make([]*models.Post, 0, 2)
	posts = append(posts, &models.Post{
		Id:        1,
		AuthorId:  1,
		Content:   "Hi There! This is Tweetster!",
		CreatedAt: time.Now(),
	})
	posts = append(posts, &models.Post{
		Id:        2,
		AuthorId:  1,
		Content:   "Good news and bad news: \\n\\n2020 is half over",
		CreatedAt: time.Now(),
	})

	return &tweetsSvc{Storage: posts}
}

func (svc tweetsSvc) PostsGetList(_ context.Context) ([]*models.Post, error) {
	return svc.Storage, nil
}

func (svc tweetsSvc) PostsGetById(_ context.Context, id uint64) (*models.Post, error) {
	for _, post := range svc.Storage {
		if post.Id == id {
			return post, nil
		}
	}

	return nil, ErrPostNotFound
}
