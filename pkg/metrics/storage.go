package metrics

import (
	"encoding/json"
	"os"
)

func SaveMetricsToFile(metric Metric) {
	file, _ := os.OpenFile("metrics.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(metric)
}
