package response

import "github.com/LordRadamanthys/teste_meli/src/adapter/output/repository"

type OrdersResponse struct {
	Order_ID string        `json:"id"`
	Itens    ItensResponse `json:"itens"`
}

type ItensResponse struct {
	ProcessedItens    []Item `json:"processed_itens"`
	NotProcessedItens []Item `json:"not_processed_itens,omitempty"`
}

type Item struct {
	ID                 string   `json:"id"`
	DistributionCenter []string `json:"distribution_center"`
}

func NewResponse(order *repository.OrdersEntity, notProcessedItens []Item,
	processedItens []Item, idOrder string) *OrdersResponse {

	for _, value := range order.Order.Itens {
		if !value.Processed {
			notProcessedItens = append(notProcessedItens, Item{
				ID: value.ID,
			})
		} else {
			processedItens = append(processedItens, Item{
				ID:                 value.ID,
				DistributionCenter: value.DistributionCenter,
			})
		}
	}

	response := &OrdersResponse{
		Order_ID: idOrder,
		Itens: ItensResponse{
			ProcessedItens:    processedItens,
			NotProcessedItens: notProcessedItens,
		},
	}
	return response
}
