package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{}

// ChangePassword godoc
// @Summary 		Change password
// @Description 	update password by user id
// @Tags 			user
// @Accept 			json
// @produce 		json
// @Param 			Authorization 	header 	string 	true 	"Authorization"
// @Param 			id path 	string 		true 	"user 	ID"
// @Param password body model.UserChangePassword true "change password"
// Success 			200  {object} 	util.Response
// Failure 			400  {object} 	util.Response
// Failure 			404  {object} 	util.Response
// Failure 			500  {object} 	util.Response
// @Router 			/users/me/password [patch]
func (u User) ChangePassword(c echo.Context) error {
	var (
		body = c.Get("body").(model.UserChangePassword)
	)
	// get id user in token
	id, err := util.GetUserId(c)
	if err != nil {
		return err

	}
	objID, _ := primitive.ObjectIDFromHex(id)

	// process
	msg, err := userService.ChangePassword(objID, body)
	if err != nil {
		return util.Response400(c, nil, msg)
	}

	return util.Response200(c, bson.M{"_id": objID}, msg)
}

// GetInfo  		godoc
// @Summary 		Get info
// @Description 	get user by id
// @Tags 			user
// @Accept 			json
// @Produce 		json
// @param 			Authorization header string true "Authorization"
// @Success      	200  {object}  util.Response
// @Failure      	400  {object}  util.Response
// @Router       	/users/me [get]
func (u User) GetInfo(c echo.Context) error {
	// Get id user in token
	id, err := util.GetUserId(c)
	if err != nil {
		return err
	}
	objID, err := primitive.ObjectIDFromHex(id)

	// process
	info, err := userService.GetInfo(objID)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)

	}

	return util.Response200(c, info, "")
}

// UpdateInfo 		godoc
// @Summary 		Update info
// @Description 	update info by userId
// @Tags 			user
// @Accept 			json
// @Produce 		json
// @param 			Authorization header string true "Authorization"
// @Param        	id path string true "user ID"
// @Param        	info  body      model.UserUpdate  true  "update info"
// @Success      	200  {object}  util.Response
// @Failure      	400  {object}  util.Response
// @Router       	/users/me [put]
func (u User) UpdateInfo(c echo.Context) error {
	var (
		body = c.Get("body").(model.UserUpdate)
	)

	// get ID user in token
	id, err := util.GetUserId(c)
	if err != nil {
		return err
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	//process
	if err := userService.UpdateInfo(objID, body); err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, bson.M{"_id": id}, "")
}
