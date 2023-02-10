package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	prefix = "aws_subnet_exporter_"
)

var (
	labels = []string{"vpcid", "subnetid", "cidrblock", "az", "name"}

	// Prometheus gauge vector for available IPs in subnets
	AvailableIPs = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: prefix + "available_ips",
		Help: "Available IPs in subnets",
	}, labels)

	// Prometheus gauge vector for max IPs in subnets
	MaxIPs = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: prefix + "max_ips",
		Help: "Max host IPs in subnet",
	}, labels)
)

// Prometheus register metrics
func RegisterMetrics() {
	prometheus.MustRegister(AvailableIPs)
	prometheus.MustRegister(MaxIPs)
}
