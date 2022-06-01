package main

import (
	"echo-app/config"
	"echo-app/dao"
	"echo-app/database"
	_ "echo-app/docs"
	"echo-app/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func init() {
	config.InitDotEnv()
	database.Connect()
	//.InitAdminAccount()
	dao.Admin{}.InitAdminAccount()
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	route.Routes(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":" + config.GetEnv().Port))

}
