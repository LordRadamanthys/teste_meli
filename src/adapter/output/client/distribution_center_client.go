package client

import (
	"errors"

	"github.com/LordRadamanthys/teste_meli/src/adapter/output/client/response"
)

type DistributionCenterClient struct {
	DistributionCenters map[string][]string
}

func NewDistributionCenterClient() *DistributionCenterClient {
	return &DistributionCenterClient{}
}

func (d *DistributionCenterClient) FindDistributionCenterByItemId(itemId string) (*response.DistributionCenterResponse, error) {

	value, ok := d.DistributionCenters[itemId]
	if !ok {
		return nil, errors.New("item not available")
	}

	return &response.DistributionCenterResponse{
		AvailableDistributionCenter: value,
	}, nil
}

// yaml
func (d *DistributionCenterClient) LoadCDs() {
	d.DistributionCenters = map[string][]string{
		"123456": {"CD1", "CD2", "CD3"},
		"123444": {"CD1", "CD2", "CD3"},
		"122212": {"CD1", "CD2", "CD3"},
		"122211": {"CD1", "CD2", "CD3"},
		"122213": {"CD1", "CD2", "CD3"},
		"122214": {"CD1", "CD2", "CD3"},
		"123457": {"CD1", "CD2", "CD3"},
	}
}
