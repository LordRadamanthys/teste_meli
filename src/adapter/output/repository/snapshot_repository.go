package repository

import (
	"encoding/json"
	"log"
	"os"
)

func (o *Orders) SaveSnapshot() error {

	if len(o.Orders) == 0 {
		return nil
	}

	o.Mu.Lock()
	defer o.Mu.Unlock()

	file, err := os.Create("orders_snapshot.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(o.Orders)
}

func (o *Orders) LoadSnapshot() error {
	o.Mu.Lock()
	defer o.Mu.Unlock()

	file, err := os.Open("orders_snapshot.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		log.Println("Error loading snapshot:", err)
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&o.Orders)
}
