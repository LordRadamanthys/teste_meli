package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/LordRadamanthys/teste_meli/src/adapter/input/controller"
	"github.com/LordRadamanthys/teste_meli/src/adapter/output/client"
	"github.com/LordRadamanthys/teste_meli/src/adapter/output/repository"
	"github.com/LordRadamanthys/teste_meli/src/application/service/order"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	if errEnv := godotenv.Load(".env"); errEnv != nil {
		log.Fatal("error: ", errEnv)
	}

	ordersRepository := repository.NewOrderRepository()
	distributionCenterClient := client.NewDistributionCenterClient()
	distributionCenterClient.LoadCDs()
	orderService := order.NewOrderService(distributionCenterClient, ordersRepository)
	orderController := controller.NewOrderController(orderService)

	r := gin.New()
	r.POST("/orders", orderController.ProcessOrder)
	r.GET("/orders/:orderId", orderController.GetOrder)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	srv := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server started on %s", os.Getenv("PORT"))

	gracefullyShutdown(srv, quit)
}

func gracefullyShutdown(srv *http.Server, quit chan os.Signal) {
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	os.Exit(0)
}
