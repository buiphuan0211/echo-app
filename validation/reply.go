package validation

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
)

type Reply struct{}

func (Reply) CreateReply(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body model.CreateReply
		if err := c.Bind(&body); err != nil {
			return util.Response400(c, nil, util.InvalidData)
		}

		if err := body.Validate(); err != nil {
			return util.Response400(c, nil, util.InvalidData)
		}

		c.Set("body", body)
		return next(c)
	}
}
