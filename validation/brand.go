package validation

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type Brand struct{}

func (b Brand) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.BrandCreateBody

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

func (b Brand) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.BrandUpdateBody

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
