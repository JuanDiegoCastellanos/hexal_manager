package handlers

//
//import (
//	"hexal_manager_v0.0.1/app/models"
//	"hexal_manager_v0.0.1/app/services"
//	"hexal_manager_v0.0.1/app/utils"
//	"net/http"
//	"strconv"
//
//	"github.com/labstack/echo/v4"
//)
//
//type AreasHandler struct {
//	Service *services.AreaService
//}
//
//func (h *AreaHandler) Create(c echo.Context) error {
//	area := new(models.Area)
//
//	if err := c.Bind(area); err != nil {
//		return err
//	}
//
//	createdArea, err := h.Service.Create(area)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusCreated, utils.Response{
//		Success: true,
//		Message: "Area created successfully",
//		Data:    createdArea,
//	})
//
//}
//
//func (h *AreaHandler) Get(c echo.Context) error {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		return err
//	}
//
//	area, err := h.Service.Get(id)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, utils.Response{
//		Success: true,
//		Data:    area,
//	})
//}
//
//func (h *AreaHandler) Update(c echo.Context) error {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		return err
//	}
//
//	area := new(models.Area)
//
//	if err := c.Bind(area); err != nil {
//		return err
//	}
//
//	updatedArea, err := h.Service.Update(id, area)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, utils.Response{
//		Success: true,
//		Message: "Area updated successfully",
//		Data:    updatedArea,
//	})
//}
//func (h *AreaHandler) Delete(c echo.Context) error {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		return err
//	}
//	err = h.Service.Delete(id)
//	if err != nil {
//		return err
//	}
//
//	return c.JSON(http.StatusOK, utils.Response{
//		Success: true,
//		Message: "Area deleted successfully",
//	})
//}
