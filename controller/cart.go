package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct{}

func (ca Cart) Create(c echo.Context) error {
	var (
		body = c.Get("body").(model.CartCreate)
	)

	// get userId
	id, err := util.GetUserId(c)
	if err != nil {
		return err
	}

	// convert to objID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// process
	cartID, err := cartService.Create(objID, body)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, cartID, util.CreateSuccessFully)
}
