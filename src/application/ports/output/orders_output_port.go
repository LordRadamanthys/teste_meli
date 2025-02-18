package output

import (
	"github.com/LordRadamanthys/teste_meli/src/adapter/output/repository"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
)

type OrdersOutputPort interface {
	SaveOrder(order domain.OrderDomain) string
	FindOrderById(orderId string) (*repository.OrdersEntity, error)
}
