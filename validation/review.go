package validation

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type Review struct{}

func (r Review) GetListReview(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query model.ReviewQuery
		if err := c.Bind(&query); err != nil {
			return util.Response400(c, nil, util.InvalidData)
		}
		c.Set("query", query)
		return next(c)
	}
}

func (r Review) CreateReview(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.CreateReview
		if err := c.Bind(&body); err != nil {
			return util.Response400(c, nil, err.Error())
		}
		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, err.Error())
		}
		c.Set("body", body)
		return next(c)
	}
}
