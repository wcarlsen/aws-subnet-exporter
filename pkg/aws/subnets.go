package aws

import (
	"context"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	errCannotParseCIDRBlock    = "cannot parse CIDR block"
	errCannotDescribeSubnets   = "cannot describe subnets"
	errUnableToCalculateMaxIPs = "unable to calculate MaxIPs from at least one CIDR block"
)

type Subnet struct {
	Name         string
	SubnetID     string
	VPCID        string
	CIDRBlock    string
	AZ           string
	AvailableIPs float64
	MaxIPs       float64
}

// Get subnets
func GetSubnets(client *ec2.Client, filter string) ([]Subnet, error) {
	log.Debug("Describing subnets")
	nameIdentifier := "tag:Name"
	resp, err := client.DescribeSubnets(context.TODO(), &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{{
			Name:   &nameIdentifier,
			Values: []string{filter},
		}},
	})
	if err != nil {
		log.Debug("Failed to describe subnets")
		return nil, errors.Wrap(err, errCannotDescribeSubnets)
	}

	var subnets []Subnet
	for _, v := range resp.Subnets {
		subnet := Subnet{
			Name:         getNameFromTags(v.Tags),
			SubnetID:     *v.SubnetId,
			VPCID:        *v.VpcId,
			CIDRBlock:    *v.CidrBlock,
			AZ:           *v.AvailabilityZone,
			AvailableIPs: float64(*v.AvailableIpAddressCount),
		}
		err := subnet.getMaxIPs()
		if err != nil {
			return nil, errors.Wrap(err, errUnableToCalculateMaxIPs)
		}
		subnets = append(subnets, subnet)
	}
	return subnets, nil
}

func (s *Subnet) getMaxIPs() error {
	_, IPNet, err := net.ParseCIDR(s.CIDRBlock)
	if err != nil {
		return errors.Wrap(err, errCannotParseCIDRBlock)
	}
	s.MaxIPs = float64(cidr.AddressCount(IPNet) - 2) // because we cannot use first and last
	return nil
}

func getNameFromTags(tags []types.Tag) string {
	for _, v := range tags {
		if *v.Key == "Name" {
			return *v.Value
		}
	}
	return "UNKNOWN"
}
