package client

import (
	"fmt"
	"testing"

	"github.com/LordRadamanthys/teste_meli/src/adapter/output/client/response"
	"github.com/stretchr/testify/assert"
)

func TestFindDistributionCenterByItemId(t *testing.T) {
	tests := []struct {
		name          string
		client        *DistributionCenterClient
		itemId        string
		expectedValue *response.DistributionCenterResponse
		expectedError error
	}{
		{
			name: "Item found",
			client: &DistributionCenterClient{
				DistributionCenters: map[string][]string{
					"123456": {"CD1", "CD3"},
				},
			},
			itemId: "123456",
			expectedValue: &response.DistributionCenterResponse{
				AvailableDistributionCenter: []string{"CD1", "CD3"},
			},
			expectedError: nil,
		},
		{
			name: "Item not found",
			client: &DistributionCenterClient{
				DistributionCenters: map[string][]string{},
			},
			itemId:        "123456",
			expectedValue: nil,
			expectedError: fmt.Errorf("item with id %s not available", "123456"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := tt.client.FindDistributionCenterByItemId(tt.itemId)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedValue, value)
			}
		})
	}
}

func TestLoadCDs(t *testing.T) {
	tests := []struct {
		name               string
		fileContent        string
		expectedCenters    map[string][]string
		expectedLogMessage string
	}{
		{
			name: "Load from file",
			fileContent: `
distribution_centers:
  "123456":
    - "CD1"
    - "CD2"
    - "CD3"
  "123444":
    - "CD1"
    - "CD2"
    - "CD3"
`,
			expectedCenters: map[string][]string{
				"123456": {"CD1", "CD2", "CD3"},
				"123444": {"CD1", "CD2", "CD3"},
			},
			expectedLogMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			client := NewDistributionCenterClient()

			client.LoadCDs("db_test.yaml")

			assert.Equal(t, tt.expectedCenters, client.DistributionCenters)
		})
	}
}
