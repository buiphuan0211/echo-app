package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct{}

func (Admin) AdminLogin(c echo.Context) error {
	var body = c.Get("body").(model.AdminLoginBody)

	// process data
	token, err := adminService.AdminLogin(body)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	// token
	data := map[string]interface{}{
		"token":   token,
		"isAdmin": true,
	}

	return util.Response200(c, data, "")
}

func (Admin) GetMyProfile(c echo.Context) error {
	// jwtPayload get id
	jwtPayload, _ := util.GetJWTPayload(c)
	// admin id
	adminID := jwtPayload["_id"].(string)
	objID, _ := primitive.ObjectIDFromHex(adminID)

	// get admin profile
	profile, err := adminService.GetMyProfile(objID)

	// if err
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	data := map[string]interface{}{
		"profile": profile,
	}

	//success
	return util.Response200(c, data, "")
}

func (Admin) UpdateMyProfile(c echo.Context) error {
	var body = c.Get("body").(model.Admin)

	// jwtPayload for get id
	jwtPayload, _ := util.GetJWTPayload(c)
	id := jwtPayload["_id"].(string)
	objID, _ := primitive.ObjectIDFromHex(id)

	// UpdateProfile
	err := adminService.UpdateMyProfile(objID, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, id, "")
}
