package client

import (
	"fmt"
	"log"
	"os"

	"github.com/LordRadamanthys/teste_meli/src/adapter/output/client/response"
	"github.com/LordRadamanthys/teste_meli/src/configuration/metrics"
	"gopkg.in/yaml.v2"
)

type DistributionCenterClient struct {
	DistributionCenters map[string][]string `yaml:"distribution_centers"`
}

func NewDistributionCenterClient() *DistributionCenterClient {
	return &DistributionCenterClient{}
}

func (d *DistributionCenterClient) FindDistributionCenterByItemId(itemId string) (*response.DistributionCenterResponse, error) {

	metrics.TotalRequestDC.Inc()

	value, ok := d.DistributionCenters[itemId]
	if !ok {
		return nil, fmt.Errorf("item with id %s not available", itemId)
	}

	return &response.DistributionCenterResponse{
		AvailableDistributionCenter: value,
	}, nil
}

func (d *DistributionCenterClient) LoadCDs() {
	data, err := os.ReadFile("./configuration/db/db.yaml")
	if err != nil {
		log.Printf("error: %s - loading CDs from memory", err.Error())
		d.DistributionCenters = loadFromMemory()
	}

	err = yaml.Unmarshal(data, &d)
	if err != nil {
		log.Printf("error: %s - loading CDs from memory", err.Error())
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
