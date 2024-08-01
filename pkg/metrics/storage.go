// pkg/metrics/storage.go
package metrics

import (
	"encoding/json"
	"log"
	"os"
)

func SaveMetricsToFile(metric Metric) {

	// Lock the metrics store to ensure thread safety

	store.Lock()
	defer store.Unlock()

	file, err := os.OpenFile("metrics.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error creating/appending/readonly metrics.json file: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(metric)
}
