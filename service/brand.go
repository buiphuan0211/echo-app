package service

import (
	"echo-app/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Brand struct{}

func (b Brand) Create(body model.BrandCreateBody) (string, error) {
	// brand BSON
	brandBSON := body.ConvertToBSON()
	// create brand
	brandID, err := brandDAO.Create(brandBSON)
	if err != nil {
		return "", err
	}

	return brandID, nil
}

func (b Brand) GetList() ([]model.BrandResponse, error) {
	var (
		listBrand = make([]model.BrandResponse, 0)
	)

	// get list brand bson
	brandsBSON, err := brandDAO.GetList()
	if err != nil {
		return listBrand, err
	}

	for _, brandBSON := range brandsBSON {
		brandJSON := brandBSON.ConvertToJSON()
		listBrand = append(listBrand, brandJSON)
	}

	return listBrand, nil
}

func (b Brand) GetByID(id primitive.ObjectID) (model.BrandResponse, error) {
	var (
		brand model.BrandResponse
	)

	// get brand by id
	brandBSON, err := brandDAO.GetByID(id)
	if err != nil {
		return brand, err
	}

	brand = brandBSON.ConvertToJSON()

	return brand, nil
}

func (b Brand) UpdateByID(id primitive.ObjectID, body model.BrandUpdateBody) (string, error) {
	return brandDAO.UpdateByID(id, body)
}

func (b Brand) DeleteByID(id primitive.ObjectID) error {
	return brandDAO.DeleteByID(id)
}
