package response

import "github.com/LordRadamanthys/teste_meli/src/adapter/output/repository"

type OrdersResponse struct {
	OrderID string        `json:"id"`
	Items   ItemsResponse `json:"items"`
}

type ItemsResponse struct {
	ProcessedItems    []Item `json:"processed_items"`
	NotProcessedItems []Item `json:"not_processed_items,omitempty"`
}

type Item struct {
	ID                        string   `json:"id"`
	PrimaryDistributionCenter string   `json:"primary_distribution_center,omitempty"`
	DistributionCenters       []string `json:"distribution_centers,omitempty"`
}

func NewResponse(order *repository.OrdersEntity, notProcessedItems []Item,
	processedItems []Item, idOrder string) *OrdersResponse {

	for _, value := range order.Order.Items {
		if !value.Processed {
			notProcessedItems = append(notProcessedItems, Item{
				ID: value.ID,
			})
		} else {
			processedItems = append(processedItems, Item{
				ID:                        value.ID,
				PrimaryDistributionCenter: value.PrimaryDistributionCenter,
				DistributionCenters:       value.DistributionCenter,
			})
		}
	}

	response := &OrdersResponse{
		OrderID: idOrder,
		Items: ItemsResponse{
			ProcessedItems:    processedItems,
			NotProcessedItems: notProcessedItems,
		},
	}
	return response
}
