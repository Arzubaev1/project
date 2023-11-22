package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// @Security ApiKeyAuth
// Create car godoc
// @ID create_car
// @Router /car [POST]
// @Summary Create Car
// @Description Create Car
// @Tags Car
// @Accept json
// @Procedure json
// @Param car body models.CreateCar true "CreateCarRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateCar(c *gin.Context) {

	var createCar models.CreateCar
	err := c.ShouldBindJSON(&createCar)
	if err != nil {
		h.handlerResponse(c, "error car should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Car().Create(c.Request.Context(), &createCar)
	if err != nil {
		h.handlerResponse(c, "storage.car.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Car().GetByID(c.Request.Context(), &models.CarPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.car.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create car resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID car godoc
// @ID get_by_id_car
// @Router /car/{id} [GET]
// @Summary Get By ID Car
// @Description Get By ID Car
// @Tags Car
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdCar(c *gin.Context) {
	var id string

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Car().GetByID(c.Request.Context(), &models.CarPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.car.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id car resposne", http.StatusOK, resp)
}

// GetList car godoc
// @ID get_list_car
// @Router /car [GET]
// @Summary Get List Car
// @Description Get List Car
// @Tags Car
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListCar(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list car offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list car limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Car().GetList(c.Request.Context(), &models.CarGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.car.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list car resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update car godoc
// @ID update_car
// @Router /car/{id} [PUT]
// @Summary Update Car
// @Description Update Car
// @Tags Car
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param car body models.UpdateCar true "UpdateCarRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateCar(c *gin.Context) {

	var (
		id        string = c.Param("id")
		updateCar models.UpdateCar
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateCar)
	if err != nil {
		h.handlerResponse(c, "error car should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateCar.Id = id
	rowsAffected, err := h.strg.Car().Update(c.Request.Context(), &updateCar)
	if err != nil {
		h.handlerResponse(c, "storage.car.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.car.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Car().GetByID(c.Request.Context(), &models.CarPrimaryKey{Id: updateCar.Id})
	if err != nil {
		h.handlerResponse(c, "storage.car.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create car resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete car godoc
// @ID delete_car
// @Router /car/{id} [DELETE]
// @Summary Delete Car
// @Description Delete Car
// @Tags Car
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteCar(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Car().Delete(c.Request.Context(), &models.CarPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.car.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create car resposne", http.StatusNoContent, nil)
}
