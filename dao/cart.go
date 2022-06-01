package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct{}

func (c Cart) Create(body model.CartCreateBSON) (string, error) {
	result, err := database.CartCol().InsertOne(util.Ctx, body)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}
