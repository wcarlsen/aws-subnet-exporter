package main

import (
	"flag"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/wcarlsen/aws-subnet-exporter/pkg/aws"
	prom "github.com/wcarlsen/aws-subnet-exporter/pkg/prometheus"
	"github.com/wcarlsen/aws-subnet-exporter/pkg/utils"
)

const (
	errGoRoutineStopped = "go routine for getting subnets stopped"
	endpoint            = "/metrics"
)

var (
	port   = flag.String("port", ":8080", "The port to listen on for HTTP requests.")
	region = flag.String("region", "eu-west-1", "AWS region")
	filter = flag.String("filter", "*eks*", "Filter subnets by tag regex when calling AWS (assumes tag key is Name")
	period = flag.Duration("period", 60*time.Second, "Period for calling AWS in seconds")
	debug  = flag.Bool("debug", false, "Enable debug logging")
)

func init() {
	utils.SetupLogger(debug)
	prom.RegisterMetrics()
}

func main() {
	log.WithFields(log.Fields{"port": *port, "region": *region, "filter": *filter, "period": *period, "endpoint": endpoint}).Info("Starting aws-subnet-exporter")
	client, err := aws.InitEC2Client(*region)
	if err != nil {
		log.Fatal(err)
	}

	cancel := make(chan struct{})

	ticker := time.NewTicker(*period)
	defer ticker.Stop()

	go func() {
		for {
			subnets, err := aws.GetSubnets(client, *filter)
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range subnets {
				prom.AvailableIPs.WithLabelValues(v.VPCID, v.SubnetID, v.CIDRBlock, v.AZ, v.Name).Set(v.AvailableIPs)
				prom.MaxIPs.WithLabelValues(v.VPCID, v.SubnetID, v.CIDRBlock, v.AZ, v.Name).Set(v.MaxIPs)
			}

			select {
			case <-ticker.C:
				continue
			case <-cancel:
				log.Fatal(errGoRoutineStopped)
			}
		}
	}()

	log.WithFields(log.Fields{"endpoint": endpoint, "port": port}).Info("Starting metrics web server")
	http.Handle(endpoint, prom.Handler)
	log.Fatal(http.ListenAndServe(*port, nil))
}
