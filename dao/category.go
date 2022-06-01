package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct{}

func (ca Category) Create(body model.CategoryBSON) (string, error) {
	// InsertOne
	result, err := database.CategoryCol().InsertOne(util.Ctx, body)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (ca Category) GetList() ([]model.CategoryBSON, error) {
	var (
		categories []model.CategoryBSON
	)

	cursor, err := database.CategoryCol().Find(util.Ctx, bson.D{})
	if err != nil {
		return categories, err
	}

	if err = cursor.All(util.Ctx, &categories); err != nil {
		return categories, nil
	}

	return categories, nil
}

func (ca Category) GetByID(id primitive.ObjectID) (model.CategoryBSON, error) {
	var (
		category model.CategoryBSON
	)

	filter := bson.M{"_id": id}

	if err := database.CategoryCol().FindOne(util.Ctx, filter).Decode(&category); err != nil {
		return category, err
	}

	return category, nil
}

func (ca Category) UpdateByID(id primitive.ObjectID, body model.CategoryUpdateBody) (string, error) {
	update := bson.M{"name": body.Name, "description": body.Description}

	// UpdateOne
	_, err := database.CategoryCol().UpdateOne(util.Ctx, bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil {
		return "", err
	}

	return "", err
}

func (ca Category) DeleteByID(id primitive.ObjectID) error {
	// filter
	filter := bson.M{"_id": id}

	// DeleteOne
	if _, err := database.CategoryCol().DeleteOne(util.Ctx, filter); err != nil {
		return err
	}

	return nil
}

func (ca Category) UpdateStatus(id primitive.ObjectID, status string) error {
	_, err := database.CategoryCol().UpdateOne(util.Ctx, bson.M{"id": id}, bson.D{
		{"$set", bson.D{{"status", status}}},
	})
	if err != nil {
		return err
	}
	return nil
}
