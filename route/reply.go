package route

import (
	"echo-app/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func reply(e *echo.Echo) {
	reviewRoutes := e.Group("/reviews", middleware.JWT([]byte(config.GetEnv().Jwt.SecretKey)))
	{
		reviewRoutes.POST("/:id", replyCtrl.CreateReply, replyVal.CreateReply)
	}

	replyRoutes := e.Group("/replies", middleware.JWT([]byte(config.GetEnv().Jwt.SecretKey)))
	{
		replyRoutes.PUT("/:id", replyCtrl.UpdateReply, replyVal.CreateReply)
		replyRoutes.DELETE("/:id", replyCtrl.DeleteReply)
	}
}
