package dao

import (
	"context"
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct{}

func (o Order) GetByUserId(id primitive.ObjectID) ([]model.OrderBSON, error) {
	var (
		orders []model.OrderBSON
	)

	pipeline := []bson.M{
		{
			"$match": bson.M{"userId": id},
		},
		{"$lookup": bson.M{
			"from":         "orderItems",
			"localField":   "items",
			"foreignField": "_id",
			"as":           "items",
		}},
	}

	cursor, err := database.OrderCol().Aggregate(util.Ctx, pipeline)
	err = cursor.All(context.Background(), &orders)
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (o Order) CreateOrder(body model.OrderCreateBSON) (string, error) {
	result, err := database.OrderCol().InsertOne(util.Ctx, body)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
