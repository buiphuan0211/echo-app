package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	ReviewBSON struct {
		ID        primitive.ObjectID `bson:"_id"`
		UserId    primitive.ObjectID `bson:"userId"`
		ProductId primitive.ObjectID `bson:"productId"`
		Rating    uint8              `bson:"rating"`
		Content   string             `bson:"content"`
		CreatedAt time.Time          `bson:"createdAt"`
		UpdatedAt time.Time          `bson:"updatedAt"`
	}

	ReviewResponse struct {
		ID        primitive.ObjectID `json:"_id"`
		UserId    primitive.ObjectID `json:"userId"`
		ProductId primitive.ObjectID `json:"productId"`
		Rating    uint8              `json:"rating"`
		Content   string             `json:"content"`
		CreatedAt time.Time          `json:"createdAt"`
		UpdatedAt time.Time          `json:"updatedAt"`
	}

	ReviewQuery struct {
		Page   int64  `query:"page"`
		Rating string `query:"rating"`
		Sort   string `query:"sort"`
	}

	CreateReview struct {
		Content string `json:"content"`
		Rating  uint8  `json:"rating"`
	}
)

func (r CreateReview) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Content, validation.Length(1, 5000).Error("length 1-5000 characters")),
	)
}
