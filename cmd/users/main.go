package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	"github.com/maxvoronov/tweetster/internal/pb"
	"github.com/maxvoronov/tweetster/internal/users/endpoints"
	"github.com/maxvoronov/tweetster/internal/users/services"
	"github.com/maxvoronov/tweetster/internal/users/transports"
)

const gRPCAddr = "127.0.0.1:8082"

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	level.Info(logger).Log("msg", "Users service started")
	defer level.Info(logger).Log("msg", "Users service ended")

	svc := services.NewUsersService()
	svc = services.LoggingMiddleware(logger)(svc)
	svcEndpoints := endpoints.PrepareServiceEndpoints(svc)
	gRPCServer := transports.NewGRPCServer(svcEndpoints)

	gRPCListener, err := net.Listen("tcp", gRPCAddr)
	if err != nil {
		level.Error(logger).Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "gRPC", "addr", gRPCAddr)
		baseServer := grpc.NewServer()
		pb.RegisterUsersServiceServer(baseServer, gRPCServer)
		baseServer.Serve(gRPCListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
