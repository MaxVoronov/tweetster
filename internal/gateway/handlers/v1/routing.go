package v1

import (
	"github.com/go-logr/logr"
	"github.com/labstack/echo/v4"

	"github.com/maxvoronov/tweetster/internal/gateway/config"
	"github.com/maxvoronov/tweetster/internal/gateway/services"
)

type Router struct {
	Config   *config.Config
	Services *services.Services
	Logger   logr.Logger
}

func NewRouter(cfg *config.Config, svc *services.Services, logger logr.Logger) *Router {
	return &Router{Config: cfg, Services: svc, Logger: logger}
}

func (r *Router) ApplyRoutes(group *echo.Group) {
	group.GET("/users/:id", r.usersGetByIDHandler)

	group.GET("/tweets", r.tweetsGetListHandler)
	group.GET("/tweets/:id", r.tweetsGetByIDHandler)
}
