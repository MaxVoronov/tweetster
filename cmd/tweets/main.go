package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/bombsimon/logrusr"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/maxvoronov/tweetster/internal/pb"
	"github.com/maxvoronov/tweetster/internal/tweets/endpoints"
	"github.com/maxvoronov/tweetster/internal/tweets/middlewares"
	"github.com/maxvoronov/tweetster/internal/tweets/services"
	"github.com/maxvoronov/tweetster/internal/tweets/transports"
	"github.com/maxvoronov/tweetster/pkg/sd/consul"
)

const gRPCAddr = "127.0.0.1"
const gRPCPort = 8801
const serviceName = "tweets-service"

func main() {
	jsonLogger := logrus.New()
	jsonLogger.SetLevel(logrus.DebugLevel)
	jsonLogger.SetFormatter(&logrus.JSONFormatter{})
	logger := logrusr.NewLogger(jsonLogger)

	svc := services.NewTweetsService()
	svc = middlewares.LoggingMiddleware(logger)(svc)
	svcEndpoints := endpoints.PrepareServiceEndpoints(svc)
	gRPCServer := transports.NewGRPCServer(svcEndpoints)

	serverAddr := net.JoinHostPort(gRPCAddr, strconv.Itoa(gRPCPort))
	gRPCListener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		logger.Error(err, "Failed to init gRPC listener", "addr", serverAddr)
		os.Exit(1)
	}

	sdRegistrar, err := consul.NewServiceRegistrar("127.0.0.1", 8500)
	if err != nil {
		logger.Error(err, "Failed to init service discovery registrar", "addr", serverAddr)
		os.Exit(1)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterTweetsServiceServer(baseServer, gRPCServer)
		if err := baseServer.Serve(gRPCListener); err != nil {
			logger.Error(err, "Failed to start gRPC server", "addr", serverAddr)
		}
	}()

	serviceID, err := sdRegistrar.Register(serviceName, gRPCAddr, gRPCPort)
	if err != nil {
		logger.Error(err, "Failed to register service in service discovery", "serviceName", serviceName)
		os.Exit(1)
	}

	logger.Info("Tweets service successfully started", "addr", serverAddr)
	logger.Error(err, "Stop service by signal", "signal", <-errs)
	if err := sdRegistrar.Deregister(serviceID); err != nil {
		logger.Error(err, "Failed to deregister service in service discovery", "serviceID", serviceID)
		os.Exit(1)
	}
	logger.Info("Tweets service successfully stopped")
}
