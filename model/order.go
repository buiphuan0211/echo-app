package model

import (
	"echo-app/util"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	OrderBSON struct {
		Id         primitive.ObjectID `bson:"_id"`
		UserId     primitive.ObjectID `bson:"userId"`
		DeliveryId primitive.ObjectID `bson:"deliveryId"`
		OrderCode  string             `bson:"orderCode"`
		Status     string             `bson:"status"`
		TotalPrice int64              `bson:"totalPrice"`
		Note       string             `bson:"note"`
		Payment    PaymentType        `bson:"payment"`
		Items      []OrderItemBSON    `bson:"items"`
		CreatedAt  time.Time          `bson:"createdAt"`
		UpdatedAt  time.Time          `bson:"updatedAt"`
	}
	OrderCreateBSON struct {
		Id         primitive.ObjectID   `bson:"_id"`
		UserId     primitive.ObjectID   `bson:"userId"`
		DeliveryId primitive.ObjectID   `bson:"deliveryId"`
		OrderCode  string               `bson:"orderCode"`
		Status     string               `bson:"status"`
		TotalPrice int64                `bson:"totalPrice"`
		Note       string               `bson:"note"`
		Payment    PaymentType          `bson:"payment"`
		Items      []primitive.ObjectID `bson:"items"`
		CreatedAt  time.Time            `bson:"createdAt"`
		UpdatedAt  time.Time            `bson:"updatedAt"`
	}

	OrderResponse struct {
		Id         primitive.ObjectID  `json:"_id"`
		UserId     primitive.ObjectID  `json:"userId"`
		DeliveryId primitive.ObjectID  `json:"deliveryId"`
		OrderCode  string              `json:"orderCode"`
		Status     string              `json:"status"`
		TotalPrice int64               `json:"totalPrice"`
		Note       string              `json:"note"`
		Payment    PaymentType         `json:"payment"`
		Items      []OrderItemResponse `json:"items"`
		CreatedAt  time.Time           `json:"createdAt"`
		UpdatedAt  time.Time           `json:"updatedAt"`
	}

	OrderCreate struct {
		UserId     primitive.ObjectID `json:"userId"`
		DeliveryId primitive.ObjectID `json:"deliveryId"`
		OrderCode  string             `json:"orderCode"`
		Status     string             `json:"status"`
		TotalPrice int64              `json:"totalPrice"`
		Note       string             `json:"note"`
		Payment    PaymentType        `json:"payment"`
		Items      []OrderItemCreate  `json:"items"`
	}
	PaymentType struct {
		Method string `json:"method" bson:"method"`
		Status bool   `json:"status" bson:"status"`
	}
)

func (u OrderCreate) ConvertToBSON(items []primitive.ObjectID, userId primitive.ObjectID) OrderCreateBSON {
	result := OrderCreateBSON{
		Id:         primitive.NewObjectID(),
		UserId:     userId,
		DeliveryId: u.DeliveryId,
		OrderCode:  u.OrderCode,
		Status:     util.OrderStatusPending,
		TotalPrice: u.TotalPrice,
		Note:       u.Note,
		Payment:    u.Payment,
		Items:      items,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	return result
}

func (o OrderBSON) ConvertToJSON(items []OrderItemResponse) OrderResponse {
	result := OrderResponse{
		Id:         o.Id,
		UserId:     o.UserId,
		DeliveryId: o.DeliveryId,
		OrderCode:  o.OrderCode,
		Status:     o.Status,
		TotalPrice: o.TotalPrice,
		Note:       o.Note,
		Payment:    o.Payment,
		Items:      items,
	}
	return result
}
