package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// @Security ApiKeyAuth
// Create order godoc
// @ID create_order
// @Router /order [POST]
// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept json
// @Procedure json
// @Param order body models.CreateOrder true "CreateOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateOrder(c *gin.Context) {

	var createOrder models.CreateOrder
	err := c.ShouldBindJSON(&createOrder)
	if err != nil {
		h.handlerResponse(c, "error order should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Order().Create(c.Request.Context(), &createOrder)
	if err != nil {
		h.handlerResponse(c, "storage.order.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Order().GetByID(c.Request.Context(), &models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create order resposne", http.StatusCreated, resp)
}

// @Security ApiKeyAuth
// GetByID order godoc
// @ID get_by_id_order
// @Router /order/{id} [GET]
// @Summary Get By ID Order
// @Description Get By ID Order
// @Tags Order
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdOrder(c *gin.Context) {
	var id string

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Order().GetByID(c.Request.Context(), &models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id order resposne", http.StatusOK, resp)
}

// GetList order godoc
// @ID get_list_order
// @Router /order [GET]
// @Summary Get List Order
// @Description Get List Order
// @Tags Order
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListOrder(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list order offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list order limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Order().GetList(c.Request.Context(), &models.OrderGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.order.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list order resposne", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// Update order godoc
// @ID update_order
// @Router /order/{id} [PUT]
// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param order body models.UpdateOrder true "UpdateOrderRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateOrder(c *gin.Context) {

	var (
		id          string = c.Param("id")
		updateOrder models.UpdateOrder
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateOrder)
	if err != nil {
		h.handlerResponse(c, "error order should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateOrder.Id = id
	rowsAffected, err := h.strg.Order().Update(c.Request.Context(), &updateOrder)
	if err != nil {
		h.handlerResponse(c, "storage.order.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.order.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Order().GetByID(c.Request.Context(), &models.OrderPrimaryKey{Id: updateOrder.Id})
	if err != nil {
		h.handlerResponse(c, "storage.order.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create order resposne", http.StatusAccepted, resp)
}

// @Security ApiKeyAuth
// Delete order godoc
// @ID delete_order
// @Router /order/{id} [DELETE]
// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteOrder(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Order().Delete(c.Request.Context(), &models.OrderPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.order.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create order resposne", http.StatusNoContent, nil)
}
