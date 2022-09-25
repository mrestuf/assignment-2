package httpserver

import (
	"assignment2/httpserver/controllers"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	order  *controllers.OrderController
}

func NewRouter(router *gin.Engine, order *controllers.OrderController) *Router {
	return &Router{
		router: router,
		order:  order,
	}
}

func (r *Router) Start(port string) {

	r.router.GET("/orders", r.order.GetAllOrders)
	r.router.POST("/orders", r.order.CreateOrder)
	r.router.DELETE("/orders/:id", r.order.DeleteOrderbyID)
	r.router.PUT("/orders/:id", r.order.UpdateOrder)
	r.router.Run(port)
}
