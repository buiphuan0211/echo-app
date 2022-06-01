package dao

import (
	"echo-app/database"
	"echo-app/model"
	"echo-app/util"
)

type CartItem struct{}

func (c CartItem) Create(body []model.CartItemBSON) error {
	var data []interface{}
	for _, t := range body {
		data = append(data, t)
	}

	if _, err := database.CartItemCol().InsertMany(util.Ctx, data); err != nil {
		return err
	}

	return nil
}
