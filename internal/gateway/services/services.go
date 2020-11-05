package services

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/maxvoronov/tweetster/internal/gateway/config"
	"github.com/maxvoronov/tweetster/internal/pb"
	"github.com/maxvoronov/tweetster/pkg/sd/consul"
)

const (
	tweetsServiceName = "tweets-service"
	userServiceName   = "users-service"

	gRPCServiceConfig = `{"loadBalancingPolicy":"round_robin"}`
)

type Services struct {
	TweetsService pb.TweetsServiceClient
	UsersService  pb.UsersServiceClient
	config        *config.Config
}

func InitServices(cfg *config.Config, logger logr.Logger) (*Services, error) {
	svc := &Services{config: cfg}
	consul.RegisterDefaultResolver(logger)

	conn, err := svc.prepareConnection(tweetsServiceName)
	if err != nil {
		return nil, err
	}
	svc.TweetsService = pb.NewTweetsServiceClient(conn)

	conn, err = svc.prepareConnection(userServiceName)
	if err != nil {
		return nil, err
	}
	svc.UsersService = pb.NewUsersServiceClient(conn)

	return svc, nil
}

func (svc *Services) prepareConnection(serviceName string) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s://%s:%d/%s", consul.Scheme, svc.config.ConsulHost, svc.config.ConsulPort, serviceName),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(3),
				grpc_retry.WithPerRetryTimeout(5*time.Millisecond),
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1*time.Millisecond)),
				grpc_retry.WithCodes(codes.ResourceExhausted, codes.Unavailable, codes.DeadlineExceeded),
			),
		),
		grpc.WithDefaultServiceConfig(gRPCServiceConfig),
	)
}
