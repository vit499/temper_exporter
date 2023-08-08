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
		Name: "temperature_of_n5101",
		Help: "Current temperature n5101",
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
	fmt.Printf("Temperature set to %f\n", temperature)
}

func main() {
	// godotenv.Load()
	// port := os.Getenv("HTTP_PORT")
	// if port == "" {
	// 	port = "9897"
	// }
	port := "9897"
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
	strPort := fmt.Sprintf("0.0.0.0:%s", port)
	fmt.Printf("Listening on %s\n", strPort)
	err := http.ListenAndServe(strPort, nil)
	if err != nil {
		fmt.Println("Error starting HTTP server:", err)
		return
	}
}
