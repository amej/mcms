/* pkg/metrics/metrics.go */
package metrics

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func GatherUsage(interval int) {
	for {
		// Collect CPU usage
		cpuPercent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("CPU Usage: %v%%\n", cpuPercent)

		// Collect Memory usage
		memStats, _ := mem.VirtualMemory()
		fmt.Printf("Memory Usage: %v%%\n", memStats.UsedPercent)

		// Wait for the next interval
		time.Sleep(time.Duration(interval) * time.Second) // Replace with configurable interval
	}
}
