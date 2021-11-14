package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.bug.st/serial"
)

type AirQualityStats struct {
	Pm10         int `json:"pm1_0"`
	Pm25         int `json:"pm2_5"`
	Pm100        int `json:"pm10"`
	Particles03  int `json:"particles_03um"`
	Particles05  int `json:"particles_05um"`
	Particles10  int `json:"particles_10um"`
	Particles25  int `json:"particles_25um"`
	Particles50  int `json:"particles_50um"`
	Particles100 int `json:"particles_100um"`
}

var pm10 = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:  "golang",
	Name:       "air_quality_stats_pm10_1m",
	Help:       "Percentiles for air quality",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	MaxAge:     60 * time.Second,
})
var pm25 = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:  "golang",
	Name:       "air_quality_stats_pm25_1m",
	Help:       "Percentiles for air quality",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	MaxAge:     60 * time.Second,
})
var pm100 = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:  "golang",
	Name:       "air_quality_stats_pm100_1m",
	Help:       "Percentiles for air quality",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	MaxAge:     60 * time.Second,
})
var particles03 = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:  "golang",
	Name:       "air_quality_stats_particles03_1m",
	Help:       "Percentiles for air quality",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	MaxAge:     60 * time.Second,
})
var particles05 = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:  "golang",
	Name:       "air_quality_stats_particles05_1m",
	Help:       "Percentiles for air quality",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	MaxAge:     60 * time.Second,
})
var particles10 = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:  "golang",
	Name:       "air_quality_stats_particles10_1m",
	Help:       "Percentiles for air quality",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	MaxAge:     60 * time.Second,
})
var particles25 = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:  "golang",
	Name:       "air_quality_stats_particles25_1m",
	Help:       "Percentiles for air quality",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	MaxAge:     60 * time.Second,
})
var particles50 = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:  "golang",
	Name:       "air_quality_stats_particles50_1m",
	Help:       "Percentiles for air quality",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	MaxAge:     60 * time.Second,
})
var particles100 = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:  "golang",
	Name:       "air_quality_stats_particles100_1m",
	Help:       "Percentiles for air quality",
	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	MaxAge:     60 * time.Second,
})

func recordAirQualityMetrics() {
	mode := &serial.Mode{
		BaudRate: 115200,
	}
	port, err := serial.Open("/dev/cu.usbserial-1410", mode)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(port)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	go func() {
		for scanner.Scan() {
			line := scanner.Bytes()
			if bytes.HasPrefix(line, []byte("{")) {
				stats := AirQualityStats{}
				json.Unmarshal(line, &stats)
				pm10.Observe(float64(stats.Pm10))
				pm25.Observe(float64(stats.Pm25))
				pm100.Observe(float64(stats.Pm100))
				particles03.Observe(float64(stats.Particles03))
				particles05.Observe(float64(stats.Particles05))
				particles10.Observe(float64(stats.Particles10))
				particles25.Observe(float64(stats.Particles25))
				particles50.Observe(float64(stats.Particles50))
				particles100.Observe(float64(stats.Particles100))
			}
			fmt.Println(string(line))
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()
}
func main() {
	fmt.Println("Hello, world.")
	recordAirQualityMetrics()
	prometheus.MustRegister(pm10)
	prometheus.MustRegister(pm25)
	prometheus.MustRegister(pm100)
	prometheus.MustRegister(particles03)
	prometheus.MustRegister(particles05)
	prometheus.MustRegister(particles10)
	prometheus.MustRegister(particles25)
	prometheus.MustRegister(particles50)
	prometheus.MustRegister(particles100)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
