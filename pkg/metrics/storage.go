package metrics

import (
	"encoding/json"
	"os"
	"time"
)

type Metric struct {
	CPUUsage    float64   `json:"cpu_usage"`
	MemoryUsage float64   `json:"memory_usage"`
	Timestamp   time.Time `json:"timestamp"`
}

func SaveMetricsToFile(metric Metric) {
	file, _ := os.OpenFile("metrics.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(metric)
}
