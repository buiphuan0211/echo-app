package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	DeliveryBSON struct {
		ID               primitive.ObjectID `bson:"_id"`
		ShippingUnitName string             `bson:"shippingUnitName"`
		ShipFee          int                `bson:"shipFee"`
		Address          string             `bson:"address"`
		CreatedAt        time.Time          `bson:"createdAt"`
		UpdatedAt        time.Time          `bson:"updatedAt"`
	}
	DeliveryResponse struct {
		ID               primitive.ObjectID `json:"_id"`
		ShippingUnitName string             `json:"shippingUnitName"`
		ShipFee          int                `json:"shipFee"`
		Address          string             `json:"address"`
		CreatedAt        time.Time          `json:"createdAt"`
		UpdatedAt        time.Time          `json:"updatedAt"`
	}

	DeliveryCreate struct {
		ShippingUnitName string `json:"shippingUnitName"`
		ShipFee          int    `json:"shipFee"`
		Address          string `json:"address"`
	}
)
