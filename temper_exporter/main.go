package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// CustomExporter is a custom Prometheus exporter.
type CustomExporter struct {
	temperatureGauge prometheus.Gauge
	mutex            sync.Mutex
}

// NewCustomExporter creates a new instance of CustomExporter.
func NewCustomExporter() *CustomExporter {
	temperatureGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "custom_temperature",
		Help: "Current temperature",
	})

	prometheus.MustRegister(temperatureGauge)

	return &CustomExporter{
		temperatureGauge: temperatureGauge,
	}
}

// SetTemperature sets the temperature gauge with a random temperature value.
func (e *CustomExporter) SetTemperature(temperature float64) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.temperatureGauge.Set(temperature)
}

func main() {
	exporter := NewCustomExporter()

	// Set temperature every 5 seconds for demonstration purposes.
	go func() {
		for {
			temperature := rand.Float64() * 100.0 // Generate a random temperature between 0 and 100
			exporter.SetTemperature(temperature)

			time.Sleep(5 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting HTTP server:", err)
		return
	}
}
