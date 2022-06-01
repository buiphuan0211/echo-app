package route

import "github.com/labstack/echo/v4"

func Route(e *echo.Echo) {
	user(e)
	admin(e)
	category(e)
	brand(e)
	auth(e)
	user(e)
	order(e)
	product(e)
	review(e)
	reply(e)
	//admin(e)
}
