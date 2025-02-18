package repository

import (
	"sync"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
)

type Orders struct {
	mu    sync.Mutex
	Order map[string]domain.OrderDomain `json:"order"`
}
