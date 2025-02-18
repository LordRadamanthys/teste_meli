package main

import (
	"github.com/LordRadamanthys/teste_meli/src/adapter/input/controller"
	"github.com/LordRadamanthys/teste_meli/src/adapter/output/client"
	"github.com/LordRadamanthys/teste_meli/src/adapter/output/repository"
	"github.com/LordRadamanthys/teste_meli/src/application/service/order"
	"github.com/gin-gonic/gin"
)

func main() {

	ordersRepository := repository.NewOrderRepository()
	distributionCenterClient := client.NewDistributionCenterClient()
	distributionCenterClient.LoadCDs()
	orderService := order.NewOrderService(distributionCenterClient, ordersRepository)
	orderController := controller.NewOrderController(orderService)

	r := gin.New()

	r.POST("/orders", orderController.ProcessOrder)
	r.GET("/orders/:orderId", orderController.GetOrder)

	r.Run(":8080")

}
