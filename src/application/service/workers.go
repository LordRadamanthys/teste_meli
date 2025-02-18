package service

import (
	"log"
	"sync"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/LordRadamanthys/teste_meli/src/application/ports/output"
)

func StartWorkers(jobsChan <-chan string, resultChan chan<- domain.ItemDomain, dc output.DistributionCenterOutputPort, wg *sync.WaitGroup) {
	numWorkers := calculateWorkers(len(jobsChan))
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
			result <- domain.ItemDomain{ID: item, DistributionCenter: []string{""}, Processed: false}
		} else {
			result <- domain.ItemDomain{ID: item, DistributionCenter: centersList.AvailableDistributionCenter, Processed: true}
		}
	}
}

func calculateWorkers(numItens int) int {
	switch {
	case numItens <= 5:
		return 1
	case numItens <= 20:
		return 2
	case numItens <= 50:
		return 5
	default:
		return 10
	}
}
