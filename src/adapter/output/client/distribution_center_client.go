package client

import (
	"errors"
	"os"

	"github.com/LordRadamanthys/teste_meli/src/adapter/output/client/response"
	"gopkg.in/yaml.v2"
)

type DistributionCenterClient struct {
	DistributionCenters map[string][]string `yaml:"distribution_centers"`
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
	data, err := os.ReadFile("./configuration/db/db.yaml")
	if err != nil {
		d.DistributionCenters = loadFromMemory()
	}

	err = yaml.Unmarshal(data, &d)
	if err != nil {
		d.DistributionCenters = loadFromMemory()
	}

}

func loadFromMemory() map[string][]string {

	return map[string][]string{
		"123456": {"CD1", "CD3"},
		"123444": {"CD3"},
		"122212": {"CD1"},
		"122211": {"CD2", "CD3"},
	}
}
