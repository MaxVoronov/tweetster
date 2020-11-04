package middleware

import (
	"time"

	"github.com/go-logr/logr"
	"github.com/labstack/echo/v4"
)

func LoggingMiddleware(logger logr.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (err error) {
			req := ctx.Request()
			res := ctx.Response()
			start := time.Now()
			if err = next(ctx); err != nil {
				ctx.Error(err)
			}
			stop := time.Now()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			if err != nil {
				logger.Error(
					err,
					req.Method+" "+req.RequestURI,
					"id", id,
					"status", res.Status,
					"latency", stop.Sub(start).Microseconds(),
					"host", req.Host,
					"protocol", req.Proto,
					"remote_ip", ctx.RealIP(),
					"referer", req.Referer(),
					"user_agent", req.UserAgent(),
				)

				return nil
			}

			logger.Info(
				req.Method+" "+req.RequestURI,
				"id", id,
				"status", res.Status,
				"latency", stop.Sub(start).Microseconds(),
				"host", req.Host,
				"protocol", req.Proto,
				"remote_ip", ctx.RealIP(),
				"referer", req.Referer(),
				"user_agent", req.UserAgent(),
			)

			return nil
		}
	}
}
