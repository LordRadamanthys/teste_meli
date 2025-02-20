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
	"github.com/LordRadamanthys/teste_meli/src/configuration/metrics"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	if errEnv := godotenv.Load(".env"); errEnv != nil {
		log.Fatal("error: ", errEnv)
	}

	ordersRepository := repository.NewOrderRepository()
	distributionCenterClient := client.NewDistributionCenterClient()
	distributionCenterClient.LoadCDs("db.yaml")
	orderService := order.NewOrderService(distributionCenterClient, ordersRepository)
	orderController := controller.NewOrderController(orderService)

	r := gin.New()
	r.Use(metrics.PrometheusMiddleware())

	r.POST("/orders", orderController.ProcessOrder)
	r.GET("/orders/:orderId", orderController.GetOrder)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	ordersRepository.LoadSnapshot()

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

	gracefullyShutdown(ctx, cancel, srv, quit, ordersRepository)
}

func gracefullyShutdown(ctx context.Context,
	cancel context.CancelFunc,
	srv *http.Server,
	quit chan os.Signal,
	ordersRepository *repository.Orders) {
	<-quit
	log.Println("Shutting down server...")

	defer cancel()

	log.Println("Waiting 5 seconds to finish processing orders...")
	time.Sleep(5 * time.Second)

	if err := ordersRepository.SaveSnapshot(); err != nil {
		log.Fatal("Error saving snapshot:", err)
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	os.Exit(0)
}
