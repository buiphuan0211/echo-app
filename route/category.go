package route

import (
	"echo-app/config"
	"echo-app/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func category(e *echo.Echo) {
	var env = config.GetEnv()
	c := e.Group("/categories")
	c.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))

	c.POST("", categoryCtrl.Create, categoryVal.CreateBody)
	c.GET("", categoryCtrl.GetList)
	c.GET("/:id", categoryCtrl.GetByID, validation.ValidateID)
	c.PUT("/:id", categoryCtrl.UpdateByID, validation.ValidateID, categoryVal.UpdateBody)
	c.DELETE("/:id", categoryCtrl.DeleteByID, validation.ValidateID)
	c.PATCH("/:id/status", categoryCtrl.UpdateStatus, validation.ValidateID)
}
