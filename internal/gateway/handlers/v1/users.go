package v1

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/maxvoronov/tweetster/internal/pb"
)

func (r *Router) usersGetByIDHandler(c echo.Context) error {
	id := c.Param("id")

	resp, err := r.Services.UsersService.UserGetByID(context.Background(), &pb.UserGetByIDRequest{Id: id})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"service":  "users",
		"method":   "get_by_id",
		"id":       id,
		"user":     resp.User,
		"app_mode": r.Config.AppMode,
	})
}
