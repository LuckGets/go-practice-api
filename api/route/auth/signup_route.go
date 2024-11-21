package authRouter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func MountAuthRoute(group *echo.Group) {
	g := group.Group("/auth")
	g.GET("/me", getMe)
	g.POST("/signup", authSignUpRouter)
}

func getMe(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}

func authSignUpRouter(c echo.Context) error {

	return nil
}
