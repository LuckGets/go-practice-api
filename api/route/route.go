package route

import (
	"github.com/LuckGets/go-echo-rest/api/handler"
	authRouter "github.com/LuckGets/go-echo-rest/api/route/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func MountMiddleWare(e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
}

func MountRoute(e *echo.Echo) {
	g := e.Group("/v1")
	// Heath check handler
	g.GET("/healthcheck", handler.HealthCheckHandler)

	// All Public APIs
	authRouter.MountAuthRoute(g)
}
