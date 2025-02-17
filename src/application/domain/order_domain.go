package domain

type OrderDomain struct {
	Itens []ItemDomain `json:"itens"`
}

func NewOrderDomain(itens []ItemDomain) *OrderDomain {
	return &OrderDomain{
		Itens: itens,
	}
}
