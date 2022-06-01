package service

import (
	"echo-app/dao"
	"echo-app/model"
	"echo-app/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var limit int64 = 10

type Product struct{}

func (p Product) GetListProduct(query model.ProductQuery) ([]model.ProductResponse, error) {
	var d = dao.Product{}
	results, err := d.GetListProduct(query)
	return results, err
}

func (p Product) GetProductDetail(ID string) (*model.ProductResponse, error) {
	ojbID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	results, err := productDAO.GetProductDetail(ojbID)
	return results, err
}

func (p Product) Create(body model.ProductCreate) (string, error) {
	var (
		productBSON model.ProductBSON
	)
	productBSON = body.ConvertToBSON()
	return productDAO.Create(productBSON)
}

func (p Product) Update(id primitive.ObjectID, body model.ProductUpdate) error {
	productBSON := body.ConvertToBSON()
	return productDAO.Update(id, productBSON)
}

func (p Product) UpdateStatus(id primitive.ObjectID) error {
	var status string
	// check status
	product, err := productDAO.GetByID(id)
	if err != nil {
		return err
	}

	if product.Status == util.ProductStatusEnabled {
		status = util.ProductStatusDisabled
	} else {
		status = util.ProductStatusEnabled
	}

	// update status
	return productDAO.UpdateStatus(id, status)
}
