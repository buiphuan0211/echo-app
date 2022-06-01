package service

import (
	"echo-app/model"
	"echo-app/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct{}

func (c Category) Create(body model.CategoryCreateBody) (string, error) {
	//convert
	categoryBSON := body.ConvertToBSON()

	// create category
	return categoryDAO.Create(categoryBSON)
}

func (c Category) GetList() ([]model.CategoryResponse, error) {
	var (
		listCategory = make([]model.CategoryResponse, 0)
	)

	// get list category bson
	categoriesBSON, err := categoryDAO.GetList()
	if err != nil {
		return listCategory, err
	}

	// convert to json db
	for _, categoryBSON := range categoriesBSON {
		categoryJSON := categoryBSON.ConvertToJSON()
		listCategory = append(listCategory, categoryJSON)
	}

	return listCategory, nil
}

func (c Category) GetByID(id primitive.ObjectID) (model.CategoryResponse, error) {
	var (
		category model.CategoryResponse
	)

	// get category by id
	categoryBSON, err := categoryDAO.GetByID(id)
	if err != nil {
		return category, err
	}

	category = categoryBSON.ConvertToJSON()
	return category, nil
}

func (c Category) UpdateByID(id primitive.ObjectID, body model.CategoryUpdateBody) (string, error) {
	return categoryDAO.UpdateByID(id, body)
}

func (c Category) DeleteByID(id primitive.ObjectID) error {
	return categoryDAO.DeleteByID(id)
}

func (c Category) UpdateStatus(id primitive.ObjectID) error {
	var status string

	// check status
	category, err := categoryDAO.GetByID(id)
	if err != nil {
		return err
	}

	if category.Status == util.CategoryStatusEnabled {
		status = util.CategoryStatusDisabled
	} else {
		status = util.CategoryStatusEnabled
	}

	// update status
	return categoryDAO.UpdateStatus(id, status)
}
