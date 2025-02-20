package repository

import (
	"testing"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderRegister(t *testing.T) {
	tests := []struct {
		name  string
		order domain.OrderDomain
	}{
		{
			name: "Register new order",
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
			orders := Orders{
				Orders: make(map[string]domain.OrderDomain),
			}
			id := orders.NewOrderRegister(tt.order)
			assert.NotEmpty(t, id, "Order ID should not be empty")
			assert.Contains(t, orders.Orders, id, "Orders map should contain the new order ID")
			assert.Equal(t, tt.order, orders.Orders[id], "The registered order should match the input order")
		})
	}
}
