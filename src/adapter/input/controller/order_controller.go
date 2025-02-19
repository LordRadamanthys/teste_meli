package controller

import (
	"fmt"
	"net/http"

	"github.com/LordRadamanthys/teste_meli/src/adapter/input/request"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/LordRadamanthys/teste_meli/src/application/ports/input"
	"github.com/LordRadamanthys/teste_meli/src/configuration/metrics"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := request.ValidateRequest(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobsChan := make(chan string, len(request.Itens))
	resultChan := make(chan domain.ItemDomain, len(request.Itens))

	orderId, err := oc.OrderService.ProcessOrder(request, jobsChan, resultChan)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	metrics.OrdersProcessGoroutinesTotal.Set(0)
	c.JSON(http.StatusCreated, gin.H{"orderId": orderId})
}

func (oc *OrderController) GetOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	response, err := oc.OrderService.FindOrder(orderId)
	if err != nil {
		if err.Error() == fmt.Sprintf("order with id %s not found", orderId) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
