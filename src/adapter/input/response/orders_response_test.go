package response

import (
	"reflect"
	"testing"

	"github.com/LordRadamanthys/teste_meli/src/adapter/output/repository"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
)

func TestNewResponse(t *testing.T) {
	tests := []struct {
		name              string
		order             *repository.OrdersEntity
		notProcessedItems []Item
		processedItems    []Item
		idOrder           string
		expectedResponse  *OrdersResponse
	}{
		{
			name: "All items processed",
			order: &repository.OrdersEntity{
				ID: "order1",
				Order: domain.OrderDomain{
					Items: []domain.ItemDomain{
						{ID: "1", Processed: true, DistributionCenter: []string{"DC1"}},
						{ID: "2", Processed: true, DistributionCenter: []string{"DC2"}},
					},
				},
			},
			notProcessedItems: []Item{},
			processedItems:    []Item{},
			idOrder:           "order1",
			expectedResponse: &OrdersResponse{
				OrderID: "order1",
				Items: ItemsResponse{
					ProcessedItems: []Item{
						{ID: "1", DistributionCenter: []string{"DC1"}},
						{ID: "2", DistributionCenter: []string{"DC2"}},
					},
					NotProcessedItems: []Item{},
				},
			},
		},
		{
			name: "Some items not processed",
			order: &repository.OrdersEntity{
				ID: "order2",
				Order: domain.OrderDomain{
					Items: []domain.ItemDomain{
						{ID: "1", Processed: false},
						{ID: "2", Processed: true, DistributionCenter: []string{"DC2"}},
					},
				},
			},
			notProcessedItems: []Item{},
			processedItems:    []Item{},
			idOrder:           "order2",
			expectedResponse: &OrdersResponse{
				OrderID: "order2",
				Items: ItemsResponse{
					ProcessedItems: []Item{
						{ID: "2", DistributionCenter: []string{"DC2"}},
					},
					NotProcessedItems: []Item{
						{ID: "1"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := NewResponse(tt.order, tt.notProcessedItems, tt.processedItems, tt.idOrder)
			if !reflect.DeepEqual(response, tt.expectedResponse) {
				t.Errorf("NewResponse() = %v, want %v", response, tt.expectedResponse)
			}
		})
	}
}
