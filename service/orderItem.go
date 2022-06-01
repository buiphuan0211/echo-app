package service

import "echo-app/model"

type OrderItem struct{}

func (o OrderItem) Create(body []model.OrderItemBSON) error {
	return orderItemDAO.Create(body)
}
