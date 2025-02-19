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
	ID                 string   `json:"id"`
	DistributionCenter []string `json:"distribution_center"`
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
				ID:                 value.ID,
				DistributionCenter: value.DistributionCenter,
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
