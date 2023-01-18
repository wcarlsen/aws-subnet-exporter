package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/pkg/errors"
)

const (
	errCannotLoadConfig = "unable to load AWS config"
)

// Initialize AWS EC2 client
func InitEC2Client(region string) (*ec2.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, errors.Wrap(err, errCannotLoadConfig)
	}
	client := ec2.NewFromConfig(cfg)

	return client, nil
}
