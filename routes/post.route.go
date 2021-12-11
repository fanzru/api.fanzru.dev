package routes
import (
	"api.fanzru.dev/controllers"
	"github.com/labstack/echo/v4"
)

func RoutePost(e *echo.Echo) *echo.Echo{
	//post
	g := e.Group("/post")
	g.GET("/", controllers.GetAllMessage)
	g.POST("/create",controllers.CreateMessage)
	g.GET("/all",controllers.Message)
	g.GET("/search",controllers.SearchMessage)
	return e
}
