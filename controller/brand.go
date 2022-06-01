package controller

import (
	"echo-app/model"
	"echo-app/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct{}

func (Brand) Create(c echo.Context) error {
	var body = c.Get("body").(model.BrandCreateBody)

	// process data
	brandID, err := brandService.Create(body)

	// if err
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, bson.M{"_id": brandID}, util.CreateSuccessFully)
}

func (Brand) GetList(c echo.Context) error {
	brands, err := brandService.GetList()
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, brands, "")
}

func (Brand) GetByID(c echo.Context) error {
	id := c.Get("id").(string)
	objID, _ := primitive.ObjectIDFromHex(id)

	// process
	brand, err := brandService.GetByID(objID)

	// if error
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, brand, "")
}

func (Brand) UpdateByID(c echo.Context) error {
	var (
		id   = c.Get("id").(string)
		body = c.Get("body").(model.BrandUpdateBody)
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	// process data
	brandID, err := brandService.UpdateByID(objID, body)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, brandID, "")

}

func (Brand) DeleteByID(c echo.Context) error {
	var id = c.Get("id").(string)

	objID, _ := primitive.ObjectIDFromHex(id)

	//process
	err := brandService.DeleteByID(objID)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}

	return util.Response200(c, nil, "")
}

func (Brand) DisabledBrand(c echo.Context) error {
	return c.JSON(http.StatusOK, "Disabled brand")
}
func (Brand) EnabledBrand(c echo.Context) error {
	return c.JSON(http.StatusOK, "Enabled brand")
}
