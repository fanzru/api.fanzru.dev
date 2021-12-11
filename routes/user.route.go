package routes

import (
	"api.fanzru.dev/controllers"
	"github.com/labstack/echo/v4"
)

func RouteUser(e *echo.Echo) *echo.Echo{
	e.GET("/", controllers.GetUser)
	e.POST("/post",controllers.Create)
	return e
}
