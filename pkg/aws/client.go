package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	errCannotLoadConfig = "unable to load AWS config"
)

// Initialize AWS EC2 client
func InitEC2Client(region string) (*ec2.Client, error) {
	log.Debug("Initializing AWS client")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Error("Failed to load AWS config")
		return nil, errors.Wrap(err, errCannotLoadConfig)
	}
	client := ec2.NewFromConfig(cfg)

	return client, nil
}
