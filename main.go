package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_prDelivery "github.com/AldiOktavianto/go-domain/module/proc/delivery/rest"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	_prDelivery.NewPrHandler(e)

	e.Logger.Info(e.Start(":" + "9090"))
}
