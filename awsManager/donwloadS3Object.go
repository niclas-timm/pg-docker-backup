package awsManager

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

func DownloadS3Object(client *s3.Client, s3FileName string, localStorageObject *os.File){

	downloader := manager.NewDownloader(client)
	_, err := downloader.Download(context.TODO(), localStorageObject, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key: aws.String(s3FileName),
	})

	if err != nil {
		panic(err.Error())
	}

	fmt.Print("Successfully downloaded backup from S3.")
}