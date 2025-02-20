package repository

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/LordRadamanthys/teste_meli/src/application/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveSnapshot(t *testing.T) {
	orders := &Orders{
		Orders: map[string]domain.OrderDomain{
			"1": {Items: []domain.ItemDomain{{ID: "item1"}}},
			"2": {Items: []domain.ItemDomain{{ID: "item2"}}},
		},
	}

	err := orders.SaveSnapshot()
	assert.NoError(t, err)

	file, err := os.Open("orders_snapshot.json")
	assert.NoError(t, err)
	defer file.Close()

	var loadedOrders map[string]domain.OrderDomain
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&loadedOrders)
	assert.NoError(t, err)

	assert.Equal(t, orders.Orders, loadedOrders)

	// Clean up
	os.Remove("orders_snapshot.json")
}

func TestLoadSnapshot(t *testing.T) {
	orders := &Orders{
		Orders: make(map[string]domain.OrderDomain),
	}

	// Create a snapshot file
	file, err := os.Create("orders_snapshot.json")
	assert.NoError(t, err)
	defer file.Close()

	snapshotData := map[string]domain.OrderDomain{
		"1": {Items: []domain.ItemDomain{{ID: "item1"}}},
		"2": {Items: []domain.ItemDomain{{ID: "item2"}}},
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(snapshotData)
	assert.NoError(t, err)

	err = orders.LoadSnapshot()
	assert.NoError(t, err)

	assert.Equal(t, snapshotData, orders.Orders)

	// Clean up
	os.Remove("orders_snapshot.json")
}
