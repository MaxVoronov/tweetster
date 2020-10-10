package services

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/maxvoronov/tweetster/internal/pb"
	"github.com/maxvoronov/tweetster/internal/tweets/models"
)

var ErrPostNotFound = status.Error(codes.NotFound, "Post not found")

type tweetsSvc struct {
	Storage      []*models.Post
	UsersService pb.UsersServiceClient
}

func NewTweetsService() TweetsService {
	posts := make([]*models.Post, 0, 2)
	posts = append(posts, &models.Post{
		ID:        1,
		AuthorID:  1,
		Content:   "Hi There! This is Tweetster!",
		CreatedAt: time.Now(),
	})
	posts = append(posts, &models.Post{
		ID:        2,
		AuthorID:  1,
		Content:   "Good news and bad news: \\n\\n2020 is half over",
		CreatedAt: time.Now(),
	})

	conn, err := grpc.Dial("127.0.0.1:8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to connect to users service: %s", err.Error())
	}

	return &tweetsSvc{
		Storage:      posts,
		UsersService: pb.NewUsersServiceClient(conn),
	}
}

func (svc tweetsSvc) PostsGetList(_ context.Context) ([]*models.Post, error) {
	// ToDo: Only for testing
	reply, err := svc.UsersService.UserGetByID(context.Background(), &pb.UserGetByIDRequest{Id: 1})
	if err != nil {
		return nil, err
	}

	log.Printf("User data: %+v", reply.User)

	return svc.Storage, nil
}

func (svc tweetsSvc) PostsGetByID(_ context.Context, id uint64) (*models.Post, error) {
	for _, post := range svc.Storage {
		if post.ID == id {
			return post, nil
		}
	}

	return nil, ErrPostNotFound
}
