package order

import (
	"fmt"
	"sync"
	"testing"

	"github.com/LordRadamanthys/teste_meli/src/adapter/input/request"
	"github.com/LordRadamanthys/teste_meli/src/adapter/output/client/response"
	"github.com/LordRadamanthys/teste_meli/src/adapter/output/repository"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/stretchr/testify/assert"
)

type MockDistributionCenterOutputPort struct{}

func (m *MockDistributionCenterOutputPort) FindDistributionCenterByItemId(itemId string) (*response.DistributionCenterResponse, error) {
	if itemId == "error" {
		return nil, fmt.Errorf("item not found")
	}
	return &response.DistributionCenterResponse{
		AvailableDistributionCenter: []string{"DC1"},
	}, nil
}

type MockOrdersOutputPort struct {
	mu     sync.Mutex
	Orders map[string]domain.OrderDomain
}

func NewMockOrdersOutputPort() *MockOrdersOutputPort {
	return &MockOrdersOutputPort{
		Orders: make(map[string]domain.OrderDomain),
	}
}

func (m *MockOrdersOutputPort) SaveOrder(order domain.OrderDomain) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	id := fmt.Sprintf("order-%d", len(m.Orders)+1)
	m.Orders[id] = order
	return id
}

func (m *MockOrdersOutputPort) FindOrderById(orderId string) (*repository.OrdersEntity, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	order, ok := m.Orders[orderId]
	if !ok {
		return nil, fmt.Errorf("order with id %s not found", orderId)
	}
	entity := &repository.OrdersEntity{
		ID:    orderId,
		Order: order,
	}
	return entity, nil
}

func TestProcessOrder(t *testing.T) {
	mockDC := &MockDistributionCenterOutputPort{}
	mockOrders := NewMockOrdersOutputPort()

	orderService := NewOrderService(mockDC, mockOrders)

	tests := []struct {
		name          string
		orderRequest  request.OrderRequest
		expectedItems []domain.ItemDomain
		expectedError error
	}{
		{
			name: "Process order successfully",
			orderRequest: request.OrderRequest{
				Items: []request.ItemRequest{
					{ID: "item1"},
				},
			},
			expectedItems: []domain.ItemDomain{
				{ID: "item1", PrimaryDistributionCenter: "DC1", DistributionCenter: []string{"DC1"}, Processed: true},
			},
			expectedError: nil,
		},
		{
			name: "Process order with error item",
			orderRequest: request.OrderRequest{
				Items: []request.ItemRequest{
					{ID: "error"},
				},
			},
			expectedItems: []domain.ItemDomain{
				{ID: "error", PrimaryDistributionCenter: "", DistributionCenter: []string{""}, Processed: false},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jobsChan := make(chan string, len(tt.orderRequest.Items))
			resultChan := make(chan domain.ItemDomain, len(tt.orderRequest.Items))

			id, err := orderService.ProcessOrder(tt.orderRequest, jobsChan, resultChan)
			assert.NoError(t, err)
			assert.NotEmpty(t, id)

			order, err := mockOrders.FindOrderById(id)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedItems, order.Order.Items)
		})
	}
}
