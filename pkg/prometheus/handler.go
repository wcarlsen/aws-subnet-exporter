package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Handler = promhttp.HandlerFor(
	prometheus.DefaultGatherer,
	promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	})
