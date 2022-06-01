package controller

import (
	"echo-app/model"
	"echo-app/util"
	"fmt"
	"github.com/labstack/echo/v4"
)

type Review struct{}

// GetListReview godoc
// @Tags review
// @Summary get list review
// Description get review by productID
// @Accept json
// @produce json
// @Param id path string true "ProductID"
// Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/products/{id}/reviews [get]
func (r Review) GetListReview(c echo.Context) error {
	ID := c.Param("id")
	query := c.Get("query").(model.ReviewQuery)
	results, err := reviewService.GetListReview(ID, query)
	if err != nil {
		fmt.Println(err.Error())
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, results, "")
}

// CreateReview godoc
// @Tags review
// @Summary create Review
// Description create review
// @Accept json
// @produce json
// @Param 		Authorization 	header 	string 	true  "Authorization"
// @Param		id path string true "Product ID"
// @Param 		product body model.CreateReview true "create review"
// Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/products/{id}/reviews [post]
func (r Review) CreateReview(c echo.Context) error {
	userId, err := util.GetUserId(c)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	productId := c.Param("id")

	body := c.Get("body").(model.CreateReview)

	result, err := reviewService.CreateReview(userId, productId, body)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, result.InsertedID, util.CreateSuccessFully)
}
