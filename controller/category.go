package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct{}

// Create  godoc
// @Tags 		category
// @Summary 	Create category
// @Description create category
// @Accept 		json
// @produce 	json
// @Param 		Authorization 	header 	string 	true 	"Authorization"
// @param category body model.CategoryCreateBody true "CategoryCreate"
// Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/categories [post]
func (ca Category) Create(c echo.Context) error {
	var body = c.Get("body").(model.CategoryCreateBody)

	// process data
	categoryID, err := categoryService.Create(body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, categoryID, "")
}

// GetList  godoc
// @Tags 		category
// @Summary 	Get categories
// @Description get list category
// @Accept 		json
// @produce 	json
// @Param 		Authorization 	header 	string 	true 	"Authorization"
//Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/categories [get]
func (ca Category) GetList(c echo.Context) error {
	categories, err := categoryService.GetList()
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, categories, "")
}

// GetByID  godoc
// @Tags 		category
// @Summary 	Get category
// @Description get category by id
// @Accept 		json
// @produce 	json
// @Param 		Authorization 	header 	string 	true 	"Authorization"
// @Param 		id 	path  string true "Category ID"
//Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/categories/{id} [get]
func (ca Category) GetByID(c echo.Context) error {
	var (
		id = c.Get("id").(string)
	)
	// to objectID
	objID, _ := primitive.ObjectIDFromHex(id)

	// process
	category, err := categoryService.GetByID(objID)
	// if error
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, category, "")
}

// UpdateByID   godoc
// @Tags 		category
// @Summary 	Update category
// @Description update category by id
// @Accept 		json
// @produce 	json
// @Param 		Authorization 	header 	string 	true 	"Authorization"
// @Param 		id 	path  string true "Category ID"
// @Param 		category body model.CategoryUpdateBody true "CategoryUpdate"
//Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/categories/{id} [put]
func (ca Category) UpdateByID(c echo.Context) error {
	var (
		id   = c.Get("id").(string)
		body = c.Get("body").(model.CategoryUpdateBody)
	)
	objID, _ := primitive.ObjectIDFromHex(id)

	// process data
	_, err := categoryService.UpdateByID(objID, body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}
	return util.Response200(c, bson.M{"_id": objID}, "")
}

// DeleteByID   godoc
// @Tags 		category
// @Summary 	Delete category
// @Description delete category by id
// @Accept 		json
// @produce 	json
// @Param 		Authorization 	header 	string 	true 	"Authorization"
// @Param 		id 	path  string true "Category ID"
//Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/categories/{id} [delete]
func (ca Category) DeleteByID(c echo.Context) error {
	var (
		id = c.Get("id").(string)
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	//process
	err := categoryService.DeleteByID(objID)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, nil, "")
}

// UpdateStatus   godoc
// @Tags 		category
// @Summary 	Update status category
// @Description update status category by id
// @Accept 		json
// @produce 	json
// @Param 		Authorization 	header 	string 	true 	"Authorization"
// @Param 		id 	path  string true "Category ID"
//Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/categories/{id}/status [patch]
func (ca Category) UpdateStatus(c echo.Context) error {
	var (
		id = c.Get("id").(string)
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	//process
	if err := categoryService.UpdateStatus(objID); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{"_id": id}, "")
}
