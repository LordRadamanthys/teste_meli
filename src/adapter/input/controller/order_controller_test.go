package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LordRadamanthys/teste_meli/src/adapter/input/request"
	"github.com/LordRadamanthys/teste_meli/src/adapter/input/response"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) ProcessOrder(req request.OrderRequest, jobsChan chan string, resultChan chan domain.ItemDomain) (string, error) {
	args := m.Called(req, jobsChan, resultChan)
	return args.String(0), args.Error(1)
}

func (m *MockOrderService) FindOrder(orderId string) (*response.OrdersResponse, error) {
	args := m.Called(orderId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*response.OrdersResponse), args.Error(1)
}

func TestProcessOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Invalid JSON input", func(t *testing.T) {
		mockOrderService := new(MockOrderService)
		orderController := NewOrderController(mockOrderService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(`invalid json`)
		c.Request, _ = http.NewRequest(http.MethodPost, "/orders", body)

		orderController.ProcessOrder(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Invalid request validation", func(t *testing.T) {
		mockOrderService := new(MockOrderService)
		orderController := NewOrderController(mockOrderService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := request.OrderRequest{}
		body, _ := json.Marshal(req)
		c.Request, _ = http.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))

		orderController.ProcessOrder(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Successful order processing", func(t *testing.T) {
		mockOrderService := new(MockOrderService)
		orderController := NewOrderController(mockOrderService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := request.OrderRequest{
			Items: []request.ItemRequest{
				{
					ID:       "123",
					Quantity: 1,
				},
			},
		}
		body, _ := json.Marshal(req)
		c.Request, _ = http.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))

		mockOrderService.On("ProcessOrder", req, mock.Anything, mock.Anything).Return("order123", nil)

		orderController.ProcessOrder(c)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockOrderService.AssertExpectations(t)
	})

	t.Run("Order processing failure", func(t *testing.T) {
		mockOrderService := new(MockOrderService)
		orderController := NewOrderController(mockOrderService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		req := request.OrderRequest{
			Items: []request.ItemRequest{
				{
					ID:       "123",
					Quantity: 1,
				},
			},
		}
		body, _ := json.Marshal(req)
		c.Request, _ = http.NewRequest(http.MethodPost, "/orders", bytes.NewBuffer(body))

		mockOrderService.On("ProcessOrder", req, mock.Anything, mock.Anything).Return("", errors.New("processing error"))

		orderController.ProcessOrder(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockOrderService.AssertExpectations(t)
	})
}
func TestGetOrder(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Order not found", func(t *testing.T) {
		mockOrderService := new(MockOrderService)
		orderController := NewOrderController(mockOrderService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		orderId := "1"
		c.Params = gin.Params{{Key: "orderId", Value: orderId}}

		mockOrderService.On("FindOrder", orderId).Return(nil, fmt.Errorf("order with id %s not found", orderId))

		orderController.GetOrder(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockOrderService.AssertExpectations(t)
	})

	t.Run("Internal server error", func(t *testing.T) {
		mockOrderService := new(MockOrderService)
		orderController := NewOrderController(mockOrderService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		orderId := "123"
		c.Params = gin.Params{{Key: "orderId", Value: orderId}}

		mockOrderService.On("FindOrder", orderId).Return(nil, errors.New("internal error"))

		orderController.GetOrder(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockOrderService.AssertExpectations(t)
	})

	t.Run("Successful order retrieval", func(t *testing.T) {
		mockOrderService := new(MockOrderService)
		orderController := NewOrderController(mockOrderService)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		orderId := "123"
		c.Params = gin.Params{{Key: "orderId", Value: orderId}}

		expectedResponse := &response.OrdersResponse{
			OrderID: orderId,
			Items:   response.ItemsResponse{},
		}

		mockOrderService.On("FindOrder", orderId).Return(expectedResponse, nil)

		orderController.GetOrder(c)

		assert.Equal(t, http.StatusOK, w.Code)
		mockOrderService.AssertExpectations(t)
	})
}
