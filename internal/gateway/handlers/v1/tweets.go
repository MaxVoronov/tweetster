package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (r *Router) tweetsGetListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"service": "tweets", "method": "get_list"})
}

func (r *Router) tweetsGetByIDHandler(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{"service": "tweets", "method": "get_by_id", "id": id})
}
