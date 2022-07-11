package awsManager

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadToS3 uploads a file to AWS S3.
func UploadToS3(c *s3.Client, filepath string){
	uploader := manager.NewUploader(c)

	file, err := os.Open(filepath)
	if err != nil {
		panic("Errow while opening file")
	}
	defer file.Close()
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("immoreport-test"),
		Key:    aws.String("backups/backup-test.sql.gz"),
		Body:   file,
	})

	if err != nil {
		fmt.Println(err.Error())
		panic("Could not upload to S3.")
	}


	fmt.Println(result.Location)

}