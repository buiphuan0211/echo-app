package service

import (
	"echo-app/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct{}

func (c Cart) Create(id primitive.ObjectID, body model.CartCreate) (string, error) {
	var (
		cartBson       model.CartCreateBSON
		cartItemsBson  = make([]model.CartItemBSON, 0)
		listIDCartItem = make([]primitive.ObjectID, 0)
	)

	// Convert CartCreate to CartBson
	// get cartItems bson
	for _, v := range body.Items {
		cartItemsBson = append(cartItemsBson, v.ConvertToBSON())
	}

	// creat many cart item
	if err := c.createCartItems(cartItemsBson); err != nil {
		return "", err
	}

	// get list id items
	for _, v := range cartItemsBson {
		listIDCartItem = append(listIDCartItem, v.Id)
	}

	// convert body cart
	cartBson = body.ConvertToBSON(listIDCartItem, id)

	// create
	cartID, err := cartDAO.Create(cartBson)
	if err != nil {
		return "", nil
	}

	return cartID, nil
}

func (c Cart) createCartItems(cartItemsBson []model.CartItemBSON) error {
	return cartItemDAO.Create(cartItemsBson)
}
