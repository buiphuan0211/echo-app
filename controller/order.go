package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct{}

// GetByUserId ListOrders 	godoc
// @Tags		order
// @Summary 	Get orders
// @Description get list order by userId
// @Accept 		json
// Produce 		json
// @param Authorization header string true "Authorization"
// @Success 	200 	{object} 	util.Response
// @Success 	400 	{object} 	util.Response
// @Success 	404 	{object} 	util.Response
// @Success 	500 	{object} 	util.Response
// @Router 		/orders [get]
func (o Order) GetByUserId(c echo.Context) error {
	// Get Id in token
	id, err := util.GetUserId(c)
	if err != nil {
		return err
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	data, err := orderService.GetByUserId(objID)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, data, "")
}

// Create godoc
// @Tags         order
// @Summary      Create
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Param payload body model.OrderCreate  true "Payload"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Router       /orders [post]
func (o Order) Create(c echo.Context) error {
	var (
		body = c.Get("body").(model.OrderCreate)
	)

	// get id user in token
	UserID, err := util.GetUserId(c)
	if err != nil {
		return err
	}

	objID, err := primitive.ObjectIDFromHex(UserID)
	if err != nil {
		return err
	}

	// process
	orderID, err := orderService.Create(objID, body)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, bson.M{"_id": orderID}, "")
}
