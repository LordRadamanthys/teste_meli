package controller

import (
	"github.com/LordRadamanthys/teste_meli/src/adapter/input/request"
	"github.com/LordRadamanthys/teste_meli/src/application/ports/input"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService input.OrderInputPort
}

func NewOrderController(orderService input.OrderInputPort) *OrderController {
	return &OrderController{
		OrderService: orderService,
	}
}

func (oc *OrderController) ProcessOrder(c *gin.Context) {
	var request request.OrderRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	response, err := oc.OrderService.ProcessOrder(request)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, response)
}

func (oc *OrderController) GetOrder(c *gin.Context) {

}
