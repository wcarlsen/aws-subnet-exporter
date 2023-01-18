package prometheus

import "github.com/prometheus/client_golang/prometheus"

// Prometheus register metrics
func RegisterMetrics() {
	prometheus.MustRegister(AvailableIPs)
	prometheus.MustRegister(MaxIPs)
}
