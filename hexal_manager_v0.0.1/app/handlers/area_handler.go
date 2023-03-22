package handlers

import (
	"github.com/labstack/echo/v4"
	"hexal_manager_v0.0.1/app/models"
	"hexal_manager_v0.0.1/app/services"
	"net/http"
	"strconv"
)

type AreaHandler struct {
	areaService *services.AreaService
}

func NewAreaHandler(areaService *services.AreaService) *AreaHandler {
	return &AreaHandler{areaService: areaService}
}
func (ah *AreaHandler) Create(c echo.Context) error {
	area := &models.Area{}
	if err := c.Bind(area); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := ah.areaService.CreateArea(area); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, area)
}

func (ah *AreaHandler) Update(c echo.Context) error {
	area := &models.Area{}
	if err := c.Bind(area); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := ah.areaService.UpdateArea(area); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, area)
}
func (ah *AreaHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	area := ah.areaService.FindAreaById(uint(id))

	if err = ah.areaService.DeleteArea(&area); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
func (ah *AreaHandler) GetAreaById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	area := ah.areaService.FindAreaById(uint(id))

	return c.JSON(http.StatusOK, area)
}
func (ah *AreaHandler) GetAllAreas(c echo.Context) error {
	areas, err := ah.areaService.FindAllArea()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, areas)
}
