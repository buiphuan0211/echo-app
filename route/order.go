package route

import (
	"echo-app/config"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func order(e *echo.Echo) {
	var env = config.GetEnv()

	o := e.Group("/orders")

	o.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	o.GET("", orderCtrl.GetByUserId)
	o.POST("", orderCtrl.Create, validation.CreateOrder)

}
