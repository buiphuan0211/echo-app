package route

import (
	"echo-app/validation"

	"github.com/labstack/echo/v4"
)

func product(e *echo.Echo) {
	var p = e.Group("/products")

	p.GET("", productCtrl.GetListProduct, productVal.GetListProduct)
	p.GET("/:id", productCtrl.GetProductDetail)
	p.POST("", productCtrl.Create, productVal.Create)
	p.PUT("/:id", productCtrl.Update, validation.ValidateID, productVal.Update)
	p.PATCH("/:id/status", productCtrl.UpdateStatus, validation.ValidateID)
}
