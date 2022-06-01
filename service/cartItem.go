package service

import (
	"echo-app/model"
)

type CartItem struct{}

func (c CartItem) Create(body []model.CartItemBSON) error {
	return cartItemDAO.Create(body)
}
