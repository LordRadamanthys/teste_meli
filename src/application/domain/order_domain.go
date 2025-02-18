package domain

import "github.com/LordRadamanthys/teste_meli/src/adapter/input/request"

type OrderDomain struct {
	Itens []ItemDomain
}

func (o *OrderDomain) NewDomain(orderReq request.OrderRequest) {
	var items []ItemDomain
	for _, item := range orderReq.Itens {
		items = append(items, ItemDomain{
			ID: item.ID,
		})
	}
	o.Itens = items
}
