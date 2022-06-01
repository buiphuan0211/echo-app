package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CartBSON struct {
		Id        primitive.ObjectID `bson:"_id"`
		UserId    primitive.ObjectID `bson:"userId"`
		Items     []CartItemBSON     `bson:"items"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	CartCreateBSON struct {
		Id        primitive.ObjectID   `bson:"_id"`
		UserId    primitive.ObjectID   `bson:"userId"`
		Items     []primitive.ObjectID `bson:"items"`
		CreatedAt time.Time            `bson:"createdAt"`
		UpdatedAt time.Time            `bson:"updatedAt"`
	}

	CartCreate struct {
		UserId primitive.ObjectID `bson:"userId"`
		Items  []CartItemsCreate  `bson:"items"`
	}

	CartResponse struct {
		Id        primitive.ObjectID `bson:"_id"`
		UserId    primitive.ObjectID `bson:"userId"`
		Items     []CartItemResponse `bson:"items"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	CartUpdate struct {
		Quantity int16 `bson:"quantity"`
	}
)

func (c CartCreate) ConvertToBSON(items []primitive.ObjectID, userId primitive.ObjectID) CartCreateBSON {
	result := CartCreateBSON{
		Id:        primitive.NewObjectID(),
		UserId:    userId,
		Items:     items,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return result
}
