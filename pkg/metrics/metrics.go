/* pkg/metrics/metrics.go */
package metrics

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// Metric represents the system usage metrics.
type Metric struct {
	Timestamp   time.Time `json:"timestamp"`
	CPUUsage    float64   `json:"cpu_usage"`
	MemoryUsage float64   `json:"memory_usage"`
}

// MetricsStore is a thread-safe store for metrics.
type MetricsStore struct {
	sync.Mutex
	metrics []Metric
}

// Global metrics store
var store = &MetricsStore{}

func GatherUsage(interval int) {
	for {
		// Collect CPU usage
		cpuPercent, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Printf("Error collecting CPU usage: %v", err)
			continue
		}
		fmt.Printf("CPU Usage: %v%%\n", cpuPercent)

		// Collect Memory usage
		memStats, err := mem.VirtualMemory()
		if err != nil {
			log.Printf("Error collecting memory usage: %v", err)
			continue
		}
		fmt.Printf("Memory Usage: %v%%\n", memStats.UsedPercent)

		// Create a new Metric instance
		metric := Metric{
			Timestamp:   time.Now(),
			CPUUsage:    cpuPercent[0],
			MemoryUsage: memStats.UsedPercent,
		}
		// Add the metric to the global store
		store.Lock()
		store.metrics = append(store.metrics, metric)
		store.Unlock()

		// Marshall metric to json
		metricJSON, err := json.Marshal(metric)
		if err != nil {
			log.Printf("Error marshalling metric to JSON: %v", err)
			continue
		}

		fmt.Println(string(metricJSON))
		go SaveMetricsToFile(metric)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
