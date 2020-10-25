package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	"github.com/maxvoronov/tweetster/internal/pb"
	"github.com/maxvoronov/tweetster/internal/users/endpoints"
	"github.com/maxvoronov/tweetster/internal/users/services"
	"github.com/maxvoronov/tweetster/internal/users/transports"
	"github.com/maxvoronov/tweetster/pkg/sd/consul"
)

const gRPCAddr = "127.0.0.1"
const gRPCPort = 8803
const serviceName = "users-service"

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

	serverAddr := net.JoinHostPort(gRPCAddr, strconv.Itoa(gRPCPort))
	gRPCListener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		level.Error(logger).Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	sdRegistrar, err := consul.NewServiceRegistrar("127.0.0.1", 8500)
	if err != nil {
		level.Error(logger).Log("during", "init service discovery registrar", "err", err)
		os.Exit(1)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "gRPC", "addr", serverAddr)
		baseServer := grpc.NewServer()
		pb.RegisterUsersServiceServer(baseServer, gRPCServer)
		baseServer.Serve(gRPCListener)
	}()

	serviceID, err := sdRegistrar.Register(serviceName, gRPCAddr, gRPCPort)
	if err != nil {
		level.Error(logger).Log("during", "service discovery register", "err", err)
		return
	}

	level.Error(logger).Log("exit", <-errs)
	if err := sdRegistrar.Deregister(serviceID); err != nil {
		level.Error(logger).Log("during", "service discovery deregister", "err", err)
	}
}
