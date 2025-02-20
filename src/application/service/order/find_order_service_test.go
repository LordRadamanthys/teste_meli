package order

import (
	"errors"
	"testing"

	"github.com/LordRadamanthys/teste_meli/src/adapter/input/response"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/stretchr/testify/assert"
)

func TestFindOrder(t *testing.T) {
	mockOrders := NewMockOrdersOutputPort()

	orderService := NewOrderService(nil, mockOrders)

	tests := []struct {
		name          string
		orderId       string
		setup         func()
		expectedOrder *response.OrdersResponse
		expectedError error
	}{
		{
			name:    "Order found",
			orderId: "order-1",
			setup: func() {
				mockOrders.SaveOrder(domain.OrderDomain{
					Items: []domain.ItemDomain{
						{ID: "item1", Processed: true, DistributionCenter: []string{"DC1"}},
						{ID: "item2", Processed: false},
					},
				})
			},
			expectedOrder: &response.OrdersResponse{
				OrderID: "order-1",
				Items: response.ItemsResponse{
					ProcessedItems: []response.Item{
						{ID: "item1", DistributionCenters: []string{"DC1"}},
					},
					NotProcessedItems: []response.Item{
						{ID: "item2"},
					},
				},
			},
			expectedError: nil,
		},
		{
			name:          "Order not found",
			orderId:       "nonexistent",
			setup:         func() {},
			expectedOrder: nil,
			expectedError: errors.New("order with id nonexistent not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			order, err := orderService.FindOrder(tt.orderId)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOrder, order)
			}
		})
	}
}
