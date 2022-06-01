package route

import (
	"echo-app/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func admin(e *echo.Echo) {
	var env = config.GetEnv()

	isLogin := middleware.JWT([]byte(env.Jwt.SecretKey))
	a := e.Group("/admin", isLogin)
	a.GET("/me", adminCtrl.GetMyProfile, isLogin)
	a.PUT("/me", adminCtrl.UpdateMyProfile, isLogin, adminVal.AdminLogin)
	//a.PATCH("/me/password", adminCtrl.ChangePasswordAdmin)
	//a.PATCH("/me/avatar", adminCtrl.ChangeAvatarAdmin)
}
