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

	file, err := os.OpenFile("metrics.json", os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Error creating/appending/readonly metrics.json file: %v", err)
		return
	}
	defer file.Close()

	// Create a struct to wrap the metrics in an array
	metricsResponse := struct {
		Metrics []Metric `json:"metrics"`
	}{
		Metrics: store.metrics,
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(metricsResponse); err != nil {
		log.Printf("Error encoding metrics to JSON: %v", err)
	}
}
