package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// @Security ApiKeyAuth
// Create driver godoc
// @ID create_driver
// @Router /driver [POST]
// @Summary Create Driver
// @Description Create Driver
// @Tags Driver
// @Accept json
// @Procedure json
// @Param driver body models.CreateDriver true "CreateDriverRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateDriver(c *gin.Context) {

	var createDriver models.CreateDriver
	err := c.ShouldBindJSON(&createDriver)
	if err != nil {
		h.handlerResponse(c, "error driver should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Driver().Create(c.Request.Context(), &createDriver)
	if err != nil {
		h.handlerResponse(c, "storage.driver.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Driver().GetByID(c.Request.Context(), &models.DriverPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.driver.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create driver resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID driver godoc
// @ID get_by_id_driver
// @Router /driver/{id} [GET]
// @Summary Get By ID Driver
// @Description Get By ID Driver
// @Tags Driver
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdDriver(c *gin.Context) {
	var id string

	// Here We Check id from Token
	val, exist := c.Get("Auth")
	if !exist {
		h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
		return
	}

	driverData := val.(helper.TokenInfo)
	if len(driverData.DriverID) > 0 {
		id = driverData.DriverID
	} else {
		id = c.Param("id")
	}

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Driver().GetByID(c.Request.Context(), &models.DriverPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.driver.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id driver resposne", http.StatusOK, resp)
}

// GetList driver godoc
// @ID get_list_driver
// @Router /driver [GET]
// @Summary Get List Driver
// @Description Get List Driver
// @Tags Driver
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListDriver(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list driver offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list driver limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Driver().GetList(c.Request.Context(), &models.DriverGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.driver.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list driver resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update driver godoc
// @ID update_driver
// @Router /driver/{id} [PUT]
// @Summary Update Driver
// @Description Update Driver
// @Tags Driver
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param driver body models.UpdateDriver true "UpdateDriverRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateDriver(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateDriver models.UpdateDriver
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateDriver)
	if err != nil {
		h.handlerResponse(c, "error driver should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateDriver.Id = id
	rowsAffected, err := h.strg.Driver().Update(c.Request.Context(), &updateDriver)
	if err != nil {
		h.handlerResponse(c, "storage.driver.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.driver.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Driver().GetByID(c.Request.Context(), &models.DriverPrimaryKey{Id: updateDriver.Id})
	if err != nil {
		h.handlerResponse(c, "storage.driver.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create driver resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete driver godoc
// @ID delete_driver
// @Router /driver/{id} [DELETE]
// @Summary Delete Driver
// @Description Delete Driver
// @Tags Driver
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteDriver(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Driver().Delete(c.Request.Context(), &models.DriverPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.driver.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create driver resposne", http.StatusNoContent, nil)
}
