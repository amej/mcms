package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Metric represents a single metric entry.
type Metric struct {
	Timestamp   time.Time `json:"timestamp"`
	CPUUsage    float64   `json:"cpu_usage"`
	MemoryUsage float64   `json:"memory_usage"`
}

// MetricsResponse represents the structure of the JSON file with a metrics array.
type MetricsResponse struct {
	Metrics []Metric `json:"metrics"`
}

// LoadMetrics loads metrics from the JSON file
func LoadMetrics() ([]Metric, error) {
	file, err := os.Open("metrics.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var metricsResponse MetricsResponse
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&metricsResponse)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Metrics are  %v", metricsResponse.Metrics)

	return metricsResponse.Metrics, nil
}

// metricsHandler handles HTTP requests for metrics.
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// Load metrics from file
	metrics, err := LoadMetrics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Start initializes the HTTP server and metrics endpoint.
func Start(address string) {
	http.HandleFunc("/metrics", metricsHandler)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err) // Handle error appropriately in production code
	}
}
