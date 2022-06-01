package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	OrderItemBSON struct {
		Id        primitive.ObjectID `bson:"_id"`
		ProductId primitive.ObjectID `bson:"productId"`
		Quantity  int16              `bson:"quantity"`
		Note      string             `bson:"note"`
		Price     int64              `bson:"price"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	OrderItemResponse struct {
		Id        primitive.ObjectID `json:"_id"`
		ProductId primitive.ObjectID `json:"productId"`
		Quantity  int16              `json:"quantity"`
		Note      string             `json:"note"`
		Price     int64              `json:"price"`
	}

	OrderItemCreate struct {
		ProductId primitive.ObjectID `json:"productId"`
		Quantity  int16              `json:"quantity"`
		Note      string             `json:"note"`
		Price     int64              `json:"price"`
	}
)

func (o OrderItemCreate) ConvertToBSON() OrderItemBSON {
	result := OrderItemBSON{
		Id:        primitive.NewObjectID(),
		ProductId: o.ProductId,
		Quantity:  o.Quantity,
		Note:      o.Note,
		Price:     o.Price,
		CreatedAt: time.Now()}
	return result
}

func (o OrderItemBSON) ConvertToJSON() OrderItemResponse {
	result := OrderItemResponse{
		Id:        o.Id,
		ProductId: o.ProductId,
		Price:     o.Price,
		Quantity:  o.Quantity,
		Note:      o.Note,
	}
	return result
}
