package main

import (
	"net"
	"net/http"
	"os"

	"github.com/bombsimon/logrusr"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"github.com/maxvoronov/tweetster/internal/gateway/config"
	apiV1 "github.com/maxvoronov/tweetster/internal/gateway/handlers/v1"
	appMiddleware "github.com/maxvoronov/tweetster/internal/gateway/middleware"
	"github.com/maxvoronov/tweetster/internal/gateway/services"
)

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

	svc, err := services.InitServices(cfg, logger)
	if err != nil {
		logger.Error(err, "Failed to init connection to services")
		os.Exit(1)
	}

	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RequestID())
	e.Use(appMiddleware.LoggingMiddleware(logger))
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	router := apiV1.NewRouter(cfg, svc, logger)
	router.ApplyRoutes(e.Group("/v1"))

	e.Logger.Fatal(e.Start(net.JoinHostPort(cfg.AppHost, cfg.AppPort)))
}
