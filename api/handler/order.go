package handler

import (
	"context"
	"fmt"
	"net/http"
	_ "rent-car/api/docs"
	"rent-car/api/models"
	"rent-car/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateOrder godoc
// @Router       /order [POST]
// @Summary      Creates a new orders
// @Description  create a new order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        car body models.Order false "order"
// @Success      201 {object} models.CreateOrder
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) CreateOrder(c *gin.Context) {
	order := models.CreateOrder{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handlerResponseLog(c,h.Log,"error while reding request", http.StatusBadRequest, err.Error())
		return
	}
	
	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	id, err := h.Services.Order().Create(ctx,order)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while creating order", http.StatusInternalServerError, err.Error())
		return
	}
	handlerResponseLog(c,h.Log,"ok", http.StatusOK, id)
}

// UpdateOrder godoc
// @Router       /order/{id} [PUT]
// @Summary      Update order
// @Description  Update order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        id path string true "order_id"
// @Param        order body models.Order true "order"
// @Success      201 {object} models.Order
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) UpdateOrder(c *gin.Context) {
	order := models.UpdateOrder{}

	if err := c.ShouldBindJSON(&order); err != nil {
		handlerResponseLog(c,h.Log,"error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	order.Id = c.Param("id")
	err := uuid.Validate(order.Id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while validating", http.StatusBadRequest, err.Error())
		return
	}

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	id, err := h.Services.Order().Update(ctx,order)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while updating customer,err", http.StatusInternalServerError, err.Error())
		return
	}
	handlerResponseLog(c,h.Log,"ok", http.StatusOK, id)
}

// GetOrderList godoc
// @Router       /orders [GET]
// @Summary      Get order list
// @Description  get order list
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201 {object} models.GetAllOrdersResponse
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetAllOrder(c *gin.Context) {
	var (
		request = models.GetAllOrdersRequest{}
	)
	request.Search = c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while parsing page", http.StatusInternalServerError, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while parsing limit", http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("page: ", page)
	fmt.Println("limit: ", limit)

	request.Page = page
	request.Limit = limit

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	orders, err := h.Services.Order().GetOrderAll(ctx,request)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while getting orders", http.StatusInternalServerError, err.Error())
		return
	}

	handlerResponseLog(c, h.Log,"ok", http.StatusOK, orders)
}

// GetOrder godoc
// @Router       /order/{id} [GET]
// @Summary      Gets order
// @Description  get order by ID
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        id path string true "order"
// @Success      201 {object} models.Order
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) GetByIDOrder(c *gin.Context) {
	id := c.Param("id")

	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	order, err := h.Services.Order().GetByIDOrder(ctx,id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while getting order by id", http.StatusInternalServerError, err.Error())
		return
	}
	handlerResponseLog(c, h.Log,"ok", http.StatusOK, order)
}

// DeleteOrder godoc
// @Router       /order/{id} [DELETE]
// @Summary      Delete order
// @Description  Delete order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        id path string true "order_id"
// @Success      201 {object} models.Response
// @Failure      400 {object} models.Response
// @Failure      404 {object} models.Response
// @Failure      500 {object} models.Response
func (h Handler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := uuid.Validate(id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while validating id", http.StatusBadRequest, err.Error())
		return
	}
	ctx,cancel:= context.WithTimeout(c,config.TimewithContex)
	defer cancel()

	err = h.Services.Customer().Delete(ctx,id)
	if err != nil {
		handlerResponseLog(c,h.Log,"error while deleting customer", http.StatusInternalServerError, err.Error())
		return
	}
	handlerResponseLog(c,h.Log,"ok", http.StatusOK, id)
}
