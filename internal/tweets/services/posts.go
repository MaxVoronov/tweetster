package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	logr "github.com/go-logr/logr"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/maxvoronov/tweetster/internal/pb"
	"github.com/maxvoronov/tweetster/internal/tweets/models"
	"github.com/maxvoronov/tweetster/pkg/sd/consul"
)

const (
	consulHost        = "127.0.0.1"
	consulPort        = 8500
	GRPCServiceConfig = `{"loadBalancingPolicy":"round_robin"}`
)

var ErrPostNotFound = status.Error(codes.NotFound, "Post not found")

type tweetsSvc struct {
	Logger       logr.Logger
	Storage      []*models.Post
	UsersService pb.UsersServiceClient
}

func NewTweetsService(logger logr.Logger) TweetsService {
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

	consul.RegisterDefaultResolver(logger)
	conn, err := grpc.DialContext(
		context.Background(),
		// consul://127.0.0.1:8500/users-service
		fmt.Sprintf("%s://%s:%d/%s", consul.Scheme, consulHost, consulPort, "users-service"),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpcRetry.UnaryClientInterceptor(
				grpcRetry.WithMax(3),
				grpcRetry.WithPerRetryTimeout(5*time.Second), // time.Millisecond
				grpcRetry.WithBackoff(grpcRetry.BackoffLinear(1*time.Millisecond)),
				grpcRetry.WithCodes(codes.ResourceExhausted, codes.Unavailable, codes.DeadlineExceeded),
			),
		),
		grpc.WithDefaultServiceConfig(GRPCServiceConfig),
	)
	if err != nil {
		logger.Error(err, "Unable to connect to users service")
		os.Exit(1)
	}

	return &tweetsSvc{
		Logger:       logger,
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
