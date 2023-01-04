package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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

	amplitude := 100.0
	frequency := 3.348
	sampleRate := 100.0
	numSamples := 30

	sine := make([]float64, numSamples)
	for i := range sine {
		sine[i] = amplitude * math.Sin(2*math.Pi*frequency*float64(i)/sampleRate)
	}

	plotWave(sine)

	for {
		// Print the modulated wave
		for i := range sine {
			value := sine[i]
			sineWave.Set(value)
			fmt.Println("Sine wave value:", i, value)
			time.Sleep(time.Second)
		}
	}
}

func waveToPlotter(wave []float64) plotter.XYs {
	pts := make(plotter.XYs, len(wave))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = wave[i]
	}
	return pts
}

func plotWave(wave []float64) {
	p := plot.New()

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err := plotutil.AddLinePoints(p,
		"sine", waveToPlotter(wave))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
