package service

import (
	"log"
	"sync"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/LordRadamanthys/teste_meli/src/application/ports/output"
	"github.com/LordRadamanthys/teste_meli/src/configuration/metrics"
)

func StartWorkers(jobsChan <-chan string, resultChan chan<- domain.ItemDomain, dc output.DistributionCenterOutputPort, wg *sync.WaitGroup) {
	numWorkers := calculateWorkers(len(jobsChan))
	metrics.OrdersProcessGoroutinesTotal.Add(float64(numWorkers))
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go WorkerCds(jobsChan, resultChan, dc, wg)
	}
}

func WorkerCds(jobs <-chan string, result chan<- domain.ItemDomain, dc output.DistributionCenterOutputPort, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range jobs {
		centersList, err := dc.FindDistributionCenterByItemId(item)
		if err != nil {
			log.Printf("warn: %s", err.Error())
			metrics.UnprocessedItemsTotal.Inc()
			result <- domain.ItemDomain{ID: item, DistributionCenter: []string{""}, Processed: false}
		} else {
			metrics.ProcessedItemsTotal.Inc()
			result <- domain.ItemDomain{ID: item, DistributionCenter: centersList.AvailableDistributionCenter, Processed: true}
		}
	}
}

func calculateWorkers(numItems int) int {
	switch {
	case numItems <= 5:
		return 1
	case numItems <= 20:
		return 2
	case numItems <= 50:
		return 5
	default:
		return 10
	}
}
