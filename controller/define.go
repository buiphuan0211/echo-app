package controller

import "echo-app/service"

var (
	authService     = service.Auth{}
	orderService    = service.Order{}
	adminService    = service.Admin{}
	categoryService = service.Category{}
	brandService    = service.Brand{}
	userService     = service.User{}
	productService  = service.Product{}
	reviewService   = service.Review{}
	replyService    = service.Reply{}
	cartService     = service.Cart{}
)
