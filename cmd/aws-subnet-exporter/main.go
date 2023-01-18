package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/wcarlsen/aws-subnet-exporter/pkg/aws"
	prom "github.com/wcarlsen/aws-subnet-exporter/pkg/prometheus"
)

var (
	port   = flag.String("port", ":8080", "The port to listen on for HTTP requests.")
	region = flag.String("region", "eu-west-1", "AWS region")
	filter = flag.String("filter", "*eks*", "Filter subnets by tag regex when calling AWS (assumes tag key is Name")
	period = flag.Duration("period", 120*time.Second, "Period for calling AWS in seconds")
)

func init() {
	prom.RegisterMetrics()
}

func main() {
	client, err := aws.InitEC2Client(*region)
	if err != nil {
		log.Fatal(err)
	}

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
			time.Sleep(*period)
		}
	}()

	http.Handle("/metrics", prom.Handler)
	log.Fatal(http.ListenAndServe(*port, nil))
}
