package main

import (
	"log"
	"metrics-collector/internal/metrics"
	"metrics-collector/internal/server"
	"time"
)

func main() {
	// Start collecting metrics in a separate goroutine
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			metrics.CollectMetrics()
			<-ticker.C
		}
	}()

	// Start the HTTP server
	if err := server.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
