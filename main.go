package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var portFlag = flag.String("p", "", "Serial port to use")
	flag.Parse()
	if *portFlag == "" {
		log.Fatalln("port is a required argument")
	}
	fmt.Println("Hello, world.")
	recordAirQualityMetrics(*portFlag)
	prometheus.MustRegister(pm10)
	prometheus.MustRegister(pm25)
	prometheus.MustRegister(pm100)
	prometheus.MustRegister(particles03)
	prometheus.MustRegister(particles05)
	prometheus.MustRegister(particles10)
	prometheus.MustRegister(particles25)
	prometheus.MustRegister(particles50)
	prometheus.MustRegister(particles100)
	http.Handle("/", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
