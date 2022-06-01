package route

import (
	"echo-app/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func review(e *echo.Echo) {
	var env = config.GetEnv()
	var p = e.Group("/products")

	p.GET("/:id/reviews", reviewCtrl.GetListReview, reviewVal.GetListReview)
	
	p.Use(middleware.JWT([]byte(env.Jwt.SecretKey)))
	p.POST("/:id/reviews", reviewCtrl.CreateReview, reviewVal.CreateReview)

}
