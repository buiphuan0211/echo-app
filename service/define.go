package service

import "echo-app/dao"

var (
	adminDAO     = dao.Admin{}
	userDAO      = dao.User{}
	orderDAO     = dao.Order{}
	orderItemDAO = dao.OrderItem{}
	categoryDAO  = dao.Category{}
	brandDAO     = dao.Brand{}
	cartDAO      = dao.Cart{}
	cartItemDAO  = dao.CartItem{}
	productDAO   = dao.Product{}
	reviewDAO    = dao.Review{}
	replyDAO     = dao.Reply{}
)
