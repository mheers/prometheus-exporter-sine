package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	// MetricName is the name of the custom metric that we are creating.
	MetricName = "sine_wave"

	// Namespace is the namespace for the custom metric.
	Namespace = "examples"

	// Subsystem is the subsystem for the custom metric.
	Subsystem = "sine_wave_modulator"
)

var (
	// sineWave is a custom metric that we are creating to represent a sine wave.
	sineWave = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: Namespace,
		Subsystem: Subsystem,
		Name:      MetricName,
		Help:      "A sine wave that is modulated by a Prometheus exporter.",
	})

	// port is the port that the Prometheus exporter will listen on.
	port = flag.Int("port", 8080, "The port that the Prometheus exporter will listen on")
)

func init() {
	// Register the custom metric with Prometheus.
	prometheus.MustRegister(sineWave)

	// Parse the command-line flags.
	flag.Parse()
}

func main() {
	// Set the sine wave to 0 initially.
	sineWave.Set(0)

	// Start the Prometheus exporter.
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
			fmt.Println(err)
		}
	}()

	// Modulate the sine wave by incrementing or decrementing its value every second.
	for i := 0; i < 100; i++ {
		sineWave.Set(math.Sin(float64(i)))
		// fmt.Println("Sine wave value:", sineWave.())
		time.Sleep(time.Second)
	}
}
