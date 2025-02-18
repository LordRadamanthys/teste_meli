package main

import (
	"github.com/LordRadamanthys/teste_meli/src/adapter/input/controller"
	"github.com/LordRadamanthys/teste_meli/src/adapter/output/client"
	"github.com/LordRadamanthys/teste_meli/src/adapter/output/repository"
	"github.com/LordRadamanthys/teste_meli/src/application/service"
	"github.com/gin-gonic/gin"
)

func main() {

	ordersRepository := repository.NewOrderRepository()
	distributionCenterClient := client.NewDistributionCenterClient()
	distributionCenterClient.LoadCDs()
	orderService := service.NewOrderService(distributionCenterClient, ordersRepository)
	orderController := controller.NewOrderController(orderService)

	r := gin.New()

	r.POST("/order/process", orderController.ProcessOrder)
	r.GET("/order/:orderId", orderController.GetOrder)

	r.Run(":8080")

}
