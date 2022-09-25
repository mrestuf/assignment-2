package controllers

import (
	"assignment2/httpserver/controllers/params"
	"assignment2/httpserver/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderController struct {
	svc *services.OrderSvc
}

func NewOrderController(svc *services.OrderSvc) *OrderController {
	return &OrderController{
		svc: svc,
	}
}

func (o *OrderController) GetAllOrders(ctx *gin.Context) {
	response := o.svc.GetAllOrders()
	WriteJsonRespnse(ctx, response)
}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var req params.OrderCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := o.svc.CreateOrder(&req)
	WriteJsonRespnse(ctx, response)
}

func (s *OrderController) UpdateOrder(ctx *gin.Context) {
	var req params.OrderUpdateRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := s.svc.UpdateOrder(id, &req)
	WriteJsonRespnse(ctx, response)
}

func (o *OrderController) DeleteOrderbyID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response := o.svc.DeleteOrderbyID(id)
	WriteJsonRespnse(ctx, response)
}
