package route

import (
	"github.com/labstack/echo/v4"
)

func auth(e *echo.Echo) {
	e.POST("/admin/admin-login", adminCtrl.AdminLogin, adminVal.AdminLogin)
}
