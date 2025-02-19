package domain

import "github.com/LordRadamanthys/teste_meli/src/adapter/input/request"

type OrderDomain struct {
	Items []ItemDomain
}

func (o *OrderDomain) NewDomain(orderReq request.OrderRequest) {
	var items []ItemDomain
	for _, item := range orderReq.Items {
		items = append(items, ItemDomain{
			ID: item.ID,
		})
	}
	o.Items = items
}
