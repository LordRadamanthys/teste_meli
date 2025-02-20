package repository

import (
	"fmt"
	"testing"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveOrder(t *testing.T) {
	tests := []struct {
		name  string
		order domain.OrderDomain
	}{
		{
			name: "Save new order",
			order: domain.OrderDomain{
				Items: []domain.ItemDomain{
					{ID: "1", Processed: true, DistributionCenter: []string{"DC1"}},
					{ID: "2", Processed: false, DistributionCenter: []string{"DC2"}},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewOrderRepository()
			id := repo.SaveOrder(tt.order)
			assert.NotEmpty(t, id, "Order ID should not be empty")
			assert.Contains(t, repo.Orders, id, "Orders map should contain the new order ID")
			assert.Equal(t, tt.order, repo.Orders[id], "The saved order should match the input order")
		})
	}
}

func TestFindOrderById(t *testing.T) {
	tests := []struct {
		name          string
		order         domain.OrderDomain
		expectedError error
	}{
		{
			name: "Find existing order",
			order: domain.OrderDomain{
				Items: []domain.ItemDomain{
					{ID: "1", Processed: true, DistributionCenter: []string{"DC1"}},
					{ID: "2", Processed: false, DistributionCenter: []string{"DC2"}},
				},
			},
			expectedError: nil,
		},
		{
			name:          "Order not found",
			order:         domain.OrderDomain{},
			expectedError: fmt.Errorf("order with id %s not found", "nonexistent"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewOrderRepository()
			id := repo.SaveOrder(tt.order)

			if tt.expectedError == nil {
				entity, err := repo.FindOrderById(id)
				assert.NoError(t, err, "Expected no error")
				assert.NotNil(t, entity, "Expected entity to be not nil")
				assert.Equal(t, id, entity.ID, "Expected entity ID to match")
				assert.Equal(t, tt.order, entity.Order, "Expected entity order to match")
			} else {
				_, err := repo.FindOrderById("nonexistent")
				assert.Error(t, err, "Expected an error")
				assert.Equal(t, tt.expectedError.Error(), err.Error(), "Expected error message to match")
			}
		})
	}
}
