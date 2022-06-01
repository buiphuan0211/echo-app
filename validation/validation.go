package validation

import (
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ValidateID ...
func ValidateID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = c.Param("id")

		// validate id
		if _, err := primitive.ObjectIDFromHex(id); err != nil {
			return util.Response400(c, nil, err.Error())
		}

		c.Set("id", id)
		return next(c)
	}
}
