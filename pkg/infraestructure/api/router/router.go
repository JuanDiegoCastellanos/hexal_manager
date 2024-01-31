package router

import (
	"hexal_manager/pkg/adapter/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/vehicles", func(ctx echo.Context) error { return c.Vehicle.GetVehicles(ctx) })
	e.POST("/vehicles", func(ctx echo.Context) error { return c.Vehicle.AddVehicle(ctx) })

	return e
}
