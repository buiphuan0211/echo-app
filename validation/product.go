package validation

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type Product struct{}

func (p Product) GetListProduct(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query model.ProductQuery
		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, util.InvalidData)
		}
		c.Set("query", query)
		return next(c)
	}
}

func (p Product) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.ProductCreate

		if err := c.Bind(&body); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("body", body)
		return next(c)
	}
}

func (p Product) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.ProductUpdate

		if err := c.Bind(&body); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("body", body)
		return next(c)
	}
}
