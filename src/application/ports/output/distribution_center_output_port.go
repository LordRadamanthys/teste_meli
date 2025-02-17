package output

import "github.com/LordRadamanthys/teste_meli/src/adapter/output/client/response"

type DistributionCenterOutputPort interface {
	FindDistributionCenterByItemId(itemId string) (*response.DistributionCenterResponse, error)
}
