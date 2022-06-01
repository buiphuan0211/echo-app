package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CartItemBSON struct {
		Id        primitive.ObjectID `bson:"_id"`
		ProductId primitive.ObjectID `bson:"productId"`
		Quantity  uint               `bson:"quantity"`
		Price     string             `bson:"price"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	CartItemResponse struct {
		Id        primitive.ObjectID `json:"_id"`
		ProductId primitive.ObjectID `json:"productId"`
		Quantity  uint               `json:"quantity"`
		Price     string             `json:"price"`
		CreatedAt time.Time          `json:"createdAt"`
		UpdatedAt time.Time          `json:"updatedAt"`
	}

	CartItemsCreate struct {
		ProductId primitive.ObjectID `json:"productId"`
		Quantity  uint               `json:"quantity"`
		Price     string             `json:"price"`
	}
)

func (c CartItemsCreate) ConvertToBSON() CartItemBSON {
	result := CartItemBSON{
		Id:        primitive.NewObjectID(),
		ProductId: c.ProductId,
		Quantity:  c.Quantity,
		Price:     c.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return result
}
