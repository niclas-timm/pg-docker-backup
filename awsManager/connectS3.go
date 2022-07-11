package awsManager

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// ConnectToS3 creates an S3 client that can be used to interact with the S3 api.
func ConnectToS3() *s3.Client{
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
	log.Fatalf("failed to load configuration, %v", err)
	}

	return s3.NewFromConfig(cfg)
}

