package internal

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/txsvc/stdlib/v2/env"
)

func StartPrometheusListener() {
	// prometheus endpoint setup
	promHost := env.GetString("prom_host", "0.0.0.0:2112")
	promMetricsPath := env.GetString("prom_metrics_path", "/metrics")

	// start the metrics listener
	go func() {
		fmt.Printf(" --> starting metrics endpoint '%s' on '%s'\n", promMetricsPath, promHost)

		http.Handle(promMetricsPath, promhttp.Handler())
		http.ListenAndServe(promHost, nil)
	}()
}
