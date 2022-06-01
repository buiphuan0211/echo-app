package controller

import (
	"echo-app/model"
	"echo-app/util"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Auth struct{}

// Register godoc
// @Tags auth
// @Summary Register
// @Description user register
// @Param register body  model.UserRegister true "user register"
// @Success      	200  {object} 	util.Response
// @Failure 		400 {object} 	util.Response
// @Failure 		404 {object} 	util.Response
// @Failure 		500 {object} 	util.Response
// @Router 			/users/register [post]
func (a Auth) Register(c echo.Context) error {
	var (
		body = c.Get("body").(model.UserRegister)
	)

	// Process data
	id, err := authService.Register(body)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// Success
	return util.Response200(c, bson.M{"_id": id}, "")
}

// Login  godoc
// @Summary Login
// @Tags auth
// @Accept       json
// @Produce      json
// @Param        payload   body    model.UserLogin  true  "login user"
// @Success      200  {object}  util.Response
// @Failure      400  {object}  util.Response
// @Failure      404  {object}  util.Response
// @Failure      500  {object}  util.Response
// @Router       /users/login [post]
func (a Auth) Login(c echo.Context) error {
	var (
		user = c.Get("body").(model.UserLogin)
	)

	// process data
	token, err := authService.Login(user)
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	data := map[string]interface{}{
		"token": token,
	}

	return util.Response200(c, data, "")
}
