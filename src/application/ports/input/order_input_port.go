package input

import (
	"github.com/LordRadamanthys/teste_meli/src/adapter/input/request"
	"github.com/LordRadamanthys/teste_meli/src/adapter/input/response"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
)

type OrderInputPort interface {
	ProcessOrder(request request.OrderRequest, jobsChan chan string, resultChan chan domain.ItemDomain) (string, error)
	FindOrder(idOrder string) (*response.OrdersResponse, error)
}
