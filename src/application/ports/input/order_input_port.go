package input

import (
	"github.com/LordRadamanthys/teste_meli/src/adapter/input/request"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
)

type OrderInputPort interface {
	ProcessOrder(request.OrderRequest) (string, error)
	FindOrder(idOrder string) (*domain.OrderDomain, error)
}
