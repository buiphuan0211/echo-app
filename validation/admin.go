package validation

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

// Admin
type Admin struct{}

// Login ...
func (a Admin) AdminLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.AdminLoginBody

		// bind request data
		if err := c.Bind(&body); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// validate
		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("body", body)
		return next(c)
	}
}

func ValidateAdminUpdateBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var admin model.Admin

		// bind request body
		if err := c.Bind(&admin); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		// success
		c.Set("body", admin)
		return next(c)
	}
}
