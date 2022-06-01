package service

import (
	"echo-app/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct{}

func (o Order) GetByUserId(id primitive.ObjectID) ([]model.OrderResponse, error) {
	var (
		orders = make([]model.OrderResponse, 0)
	)

	// get orders in db
	ordersBSON, err := orderDAO.GetByUserId(id)
	if err != nil {
		return orders, err
	}

	// Convert to OrderJSON
	for _, orderBSON := range ordersBSON {
		orderItems := make([]model.OrderItemResponse, 0)

		// convert items in order
		for _, v := range orderBSON.Items {
			orderItem := v.ConvertToJSON()
			orderItems = append(orderItems, orderItem)
		}

		orderJSON := orderBSON.ConvertToJSON(orderItems)
		orders = append(orders, orderJSON)
	}

	return orders, nil
}

func (o Order) Create(id primitive.ObjectID, body model.OrderCreate) (string, error) {
	var (
		orderBSON      model.OrderCreateBSON
		listItemJson   = make([]model.OrderItemBSON, 0)
		ListIDItemJson = make([]primitive.ObjectID, 0)
	)

	// insert to order-items db

	for _, v := range body.Items {
		listItemJson = append(listItemJson, v.ConvertToBSON())
	}

	if err := o.createItems(listItemJson); err != nil {
		return "", err
	}

	// get list id insert to order db

	for _, v := range listItemJson {
		ListIDItemJson = append(ListIDItemJson, v.Id)
	}

	// convert orderCreate to orderCreateBson
	orderBSON = body.ConvertToBSON(ListIDItemJson, id)

	// create
	orderID, err := orderDAO.CreateOrder(orderBSON)
	if err != nil {
		return "", err
	}

	return orderID, nil
}

func (o Order) createItems(body []model.OrderItemBSON) error {
	var oi OrderItem
	return oi.Create(body)
}
