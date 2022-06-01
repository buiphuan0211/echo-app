package controller

import (
	"echo-app/model"
	"echo-app/util"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct{}

// GetListProduct godoc
// @Tags product
// @Summary 	Get products
// @Description get list product
// @Accept 		json
// @produce 	json
// Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/products [get]
func (p Product) GetListProduct(c echo.Context) error {
	//get query from middleware
	query := c.Get("query").(model.ProductQuery)

	results, err := productService.GetListProduct(query)
	if err != nil {
		fmt.Println(err.Error())
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, results, "")
}

// GetProductDetail  godoc
// @Tags product
// @Summary 	Get product
// @Description get product by id
// @Accept 		json
// @produce 	json
// Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/products/{id} [get]
func (p Product) GetProductDetail(c echo.Context) error {
	id := c.Param("id")

	results, err := productService.GetProductDetail(id)
	if err != nil {
		return util.Response400(c, nil, util.InvalidData)
	}
	return util.Response200(c, results, "")
}

// Create  godoc
// @Tags 		product
// @Summary 	Create product
// @Description create product
// @Accept 		json
// @produce 	json
// @Param 		Authorization 	header 	string 	true 	"Authorization"
// @param product body model.ProductCreate true "ProductCreate"
// Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/products [post]
func (p Product) Create(c echo.Context) (err error) {
	var (
		body = c.Get("body").(model.ProductCreate)
	)

	productID, err := productService.Create(body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{"_id": productID}, "")
}

// Update  godoc
// @Tags 		product
// @Summary 	Update product
// @Description update product
// @Accept 		json
// @produce 	json
// @Param 		Authorization 	header 	string 	true 	"Authorization"
// @Param		id path string true "Product ID"
// @Param 		product body model.ProductUpdate true "UpdateProduct"
// Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/products/{id} [put]
func (p Product) Update(c echo.Context) (err error) {
	var (
		body = c.Get("body").(model.ProductUpdate)
		id   = c.Get("id").(string)
	)

	objID, _ := primitive.ObjectIDFromHex(id)

	if err := productService.Update(objID, body); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{"_id": id}, "")
}

// UpdateStatus  godoc
// @Tags 		product
// @Summary 	Update status product
// @Description update status product
// @Accept 		json
// @produce 	json
// @Param 		Authorization 	header 	string 	true 	"Authorization"
// @Param 		id path string true "Product ID"
// Success 		200 {object}  util.Response
// Failure 		400 {object}  util.Response
// Failure 		404 {object}  util.Response
// Failure 		500 {object}  util.Response
// @Router 		/products/{id} [patch]
func (p Product) UpdateStatus(c echo.Context) error {
	var (
		id = c.Get("id").(string)
	)

	objID, _ := primitive.ObjectIDFromHex(id)
	//process
	if err := productService.UpdateStatus(objID); err != nil {
		return util.Response400(c, nil, err.Error())
	}

	return util.Response200(c, bson.M{"_id": id}, "")
}
