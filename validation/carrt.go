package validation

import (
	"echo-app/model"

	"github.com/labstack/echo/v4"
)

type Cart struct{}

func (c Cart) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.CartCreate

		if err := c.Bind(&body); err != nil {
			return err
		}

		// Success
		c.Set("body", body)
		return next(c)
	}
}
