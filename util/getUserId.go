package util

import (
	"github.com/labstack/echo/v4"
)

func GetUserId(c echo.Context) (string, error) {
	// GetJWTPayload
	jwtPayload, err := GetJWTPayload(c)
	if err != nil {
		return "", err
	}

	id, ok := jwtPayload["_id"].(string)
	if !ok {
		return "", err
	}
	return id, nil
}
