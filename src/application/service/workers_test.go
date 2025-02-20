package service

import (
	"fmt"
	"sync"
	"testing"

	"github.com/LordRadamanthys/teste_meli/src/adapter/output/client/response"
	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/stretchr/testify/assert"
)

type MockDistributionCenterOutputPort struct{}

func (m *MockDistributionCenterOutputPort) FindDistributionCenterByItemId(itemId string) (*response.DistributionCenterResponse, error) {
	if itemId == "error" {
		return nil, fmt.Errorf("item not found")
	}
	return &response.DistributionCenterResponse{
		AvailableDistributionCenter: []string{"DC1", "DC2"},
	}, nil
}

func TestCalculateWorkers(t *testing.T) {
	tests := []struct {
		numItems int
		expected int
	}{
		{numItems: 3, expected: 1},
		{numItems: 10, expected: 2},
		{numItems: 30, expected: 5},
		{numItems: 100, expected: 10},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("numItems=%d", tt.numItems), func(t *testing.T) {
			result := calculateWorkers(tt.numItems)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestWorkerCds(t *testing.T) {
	jobs := make(chan string, 2)
	results := make(chan domain.ItemDomain, 2)
	var wg sync.WaitGroup

	mockDC := &MockDistributionCenterOutputPort{}

	jobs <- "item1"
	jobs <- "error"
	close(jobs)

	wg.Add(1)
	go WorkerCds(jobs, results, mockDC, &wg)
	wg.Wait()
	close(results)

	expectedResults := []domain.ItemDomain{
		{ID: "item1", DistributionCenter: []string{"DC1", "DC2"}, Processed: true},
		{ID: "error", DistributionCenter: []string{""}, Processed: false},
	}

	var actualResults []domain.ItemDomain
	for result := range results {
		actualResults = append(actualResults, result)
	}

	assert.ElementsMatch(t, expectedResults, actualResults)
}

func TestStartWorkers(t *testing.T) {
	jobs := make(chan string, 10)
	results := make(chan domain.ItemDomain, 10)
	var wg sync.WaitGroup

	mockDC := &MockDistributionCenterOutputPort{}

	for i := 0; i < 10; i++ {
		jobs <- fmt.Sprintf("item%d", i)
	}
	close(jobs)

	StartWorkers(jobs, results, mockDC, &wg)
	wg.Wait()
	close(results)

	assert.Equal(t, 10, len(results))
}
