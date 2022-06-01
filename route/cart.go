package route

import (
	"echo-app/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func cart(e *echo.Echo) {
	var (
		env  = config.GetEnv()
		cart = e.Group("/cart")
	)

	cart.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	cart.POST("", cartCtrl.Create, cartVal.Create)

}
