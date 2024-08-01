package server

import (
	"bufio"
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

// LoadMetrics loads metrics from the JSON file, where each line is a separate JSON object.
func LoadMetrics() ([]Metric, error) {
	file, err := os.Open("metrics.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var metrics []Metric
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var metric Metric
		line := scanner.Text()
		if err := json.Unmarshal([]byte(line), &metric); err != nil {
			return nil, fmt.Errorf("failed to unmarshal line: %v", err)
		}
		metrics = append(metrics, metric)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %v", err)
	}

	return metrics, nil
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

// filterMetricsByTimeRange filters metrics to include only those within the specified time range.
func filterMetricsByTimeRange(metrics []Metric, startTime, endTime time.Time) []Metric {
	var filteredMetrics []Metric
	for _, metric := range metrics {
		if metric.Timestamp.After(startTime) && metric.Timestamp.Before(endTime) {
			filteredMetrics = append(filteredMetrics, metric)
		}
	}
	return filteredMetrics
}

// averageCPUUsage calculates the average CPU usage from a slice of metrics.
func averageCPUUsage(metrics []Metric) float64 {
	if len(metrics) == 0 {
		return 0
	}

	var totalCPUUsage float64
	for _, metric := range metrics {
		totalCPUUsage += metric.CPUUsage
	}
	return totalCPUUsage / float64(len(metrics))
}

/* TODO:  Fix timeFormating issue
// metricsInRangeHandler handles HTTP requests to retrieve metrics for a specific time range.
func metricsInRangeHandler(w http.ResponseWriter, r *http.Request) {
	metrics, err := LoadMetrics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	startTimeStr := r.URL.Query().Get("start")
	endTimeStr := r.URL.Query().Get("end")
	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		http.Error(w, "Invalid start time format", http.StatusBadRequest)
		return
	}
	endTime, err := time.Parse(time.RFC3339, endTimeStr)
	if err != nil {
		http.Error(w, "Invalid end time format", http.StatusBadRequest)
		return
	}

	filteredMetrics := filterMetricsByTimeRange(metrics, startTime, endTime)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(filteredMetrics); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
*/

// aggregateMetricsHandler handles HTTP requests to aggregate metrics, e.g., average CPU usage over the last hour.
func aggregateMetricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics, err := LoadMetrics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now()
	lastHourStart := now.Add(-1 * time.Hour)
	filteredMetrics := filterMetricsByTimeRange(metrics, lastHourStart, now)

	avgCPUUsage := averageCPUUsage(filteredMetrics)
	response := map[string]float64{"average_cpu_usage_last_hour": avgCPUUsage}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Start initializes the HTTP server and metrics endpoint.
func Start(address string) {
	http.HandleFunc("/metrics", metricsHandler)
	/* TODO: ttp.HandleFunc("/metrics/range", metricsInRangeHandler) */
	http.HandleFunc("/metrics/aggregate", aggregateMetricsHandler)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err) // TODO: Move to gin or other frameworks.
	}
}
