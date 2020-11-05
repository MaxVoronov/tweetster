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
	"github.com/maxvoronov/tweetster/internal/users/config"
	"github.com/maxvoronov/tweetster/internal/users/endpoints"
	"github.com/maxvoronov/tweetster/internal/users/middlewares"
	"github.com/maxvoronov/tweetster/internal/users/services"
	"github.com/maxvoronov/tweetster/internal/users/transports"
	"github.com/maxvoronov/tweetster/pkg/sd/consul"
)

const serviceName = "users-service"

func main() {
	jsonLogger := logrus.New()
	jsonLogger.SetLevel(logrus.DebugLevel)
	jsonLogger.SetFormatter(&logrus.JSONFormatter{})
	logger := logrusr.NewLogger(jsonLogger)

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error(err, "Failed to load config")
		os.Exit(1)
	}

	svc := services.NewUsersService()
	svc = middlewares.LoggingMiddleware(logger)(svc)
	svcEndpoints := endpoints.PrepareServiceEndpoints(svc)
	gRPCServer := transports.NewGRPCServer(svcEndpoints)

	serverAddr := net.JoinHostPort(cfg.AppHost, strconv.Itoa(cfg.AppPort))
	gRPCListener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		logger.Error(err, "Failed to init gRPC listener", "addr", serverAddr)
		os.Exit(1)
	}

	sdRegistrar, err := consul.NewServiceRegistrar(cfg.ConsulHost, cfg.ConsulPort)
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
		pb.RegisterUsersServiceServer(baseServer, gRPCServer)
		if err := baseServer.Serve(gRPCListener); err != nil {
			logger.Error(err, "Failed to start gRPC server", "addr", serverAddr)
		}
	}()

	serviceID, err := sdRegistrar.Register(serviceName, cfg.AppHost, cfg.AppPort)
	if err != nil {
		logger.Error(err, "Failed to register service in service discovery", "serviceName", serviceName)
		os.Exit(1)
	}

	logger.Info("Users service successfully started", "addr", serverAddr)
	logger.Error(err, "Stop service by signal", "signal", <-errs)
	if err := sdRegistrar.Deregister(serviceID); err != nil {
		logger.Error(err, "Failed to deregister service in service discovery", "serviceID", serviceID)
		os.Exit(1)
	}
	logger.Info("Users service successfully stopped")
}
