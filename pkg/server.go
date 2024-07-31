package cmd

import (
	"encoding/json"
	"net/http"
	"time"
)

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	// Load metrics from file or database
	// For simplicity, returning static data
	metrics := []Metric{
		{CPUUsage: 50.0, MemoryUsage: 30.0, Timestamp: time.Now()},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func Start(addr string) {
	http.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(addr, nil)
}
