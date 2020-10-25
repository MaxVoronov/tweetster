package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bombsimon/logrusr"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/maxvoronov/tweetster/internal/pb"
	"github.com/maxvoronov/tweetster/pkg/sd/consul"
)

const (
	consulHost        = "127.0.0.1"
	consulPort        = 8500
	GRPCServiceConfig = `{"loadBalancingPolicy":"round_robin"}`
)

func main() {
	jsonLogger := logrus.New()
	jsonLogger.SetLevel(logrus.DebugLevel)
	jsonLogger.SetFormatter(&logrus.JSONFormatter{})
	logger := logrusr.NewLogger(jsonLogger)

	consul.RegisterDefaultResolver(logger)

	conn, err := grpc.DialContext(
		context.Background(),
		// consul://127.0.0.1:8500/users-service
		fmt.Sprintf("%s://%s:%d/%s", consul.Scheme, consulHost, consulPort, "users-service"),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(3),
				grpc_retry.WithPerRetryTimeout(5*time.Second), // time.Millisecond
				grpc_retry.WithBackoff(grpc_retry.BackoffLinear(1*time.Millisecond)),
				grpc_retry.WithCodes(codes.ResourceExhausted, codes.Unavailable, codes.DeadlineExceeded),
			),
		),
		grpc.WithDefaultServiceConfig(GRPCServiceConfig),
	)
	if err != nil {
		logger.Error(err, "Unable to connect to users service")
		os.Exit(1)
	}

	userService := pb.NewUsersServiceClient(conn)
	resp, err := userService.UserGetByID(context.Background(), &pb.UserGetByIDRequest{Id: 1})
	if err != nil {
		logger.Error(err, "Failed call to UserGetByID")
		os.Exit(1)
	}

	logger.Info("Success!", "user", resp.User)
}
