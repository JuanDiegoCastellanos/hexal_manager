package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"hexal_manager_v0.0.1/app/handlers"
	"hexal_manager_v0.0.1/app/repositories"
	"hexal_manager_v0.0.1/app/services"
)

func RegisterAreaRoutes(e *echo.Echo, db *gorm.DB) {
	areaRepo := repositories.NewAreaRepository(db)
	areaService := services.NewAreaService(areaRepo)
	areaHandler := handlers.NewAreaHandler(areaService)

	// Middleware for authorization
	//e.Use(middlewares.)

	// Routes for areas
	e.POST("api_hexal/areas", areaHandler.Create)
	e.PUT("api_hexal/areas", areaHandler.Update)
	e.DELETE("api_hexal/areas/:id", areaHandler.Delete)
	e.GET("api_hexal/areas/:id", areaHandler.GetAreaById)
	e.GET("api_hexal/areas", areaHandler.GetAllAreas)

}
