package awsManager

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/NiclasTimmeDev/pg-docker-backup/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadToS3 uploads a file to AWS S3.
func UploadToS3(c *s3.Client, filepath string, config config.Config){
	uploader := manager.NewUploader(c)

	file, err := os.Open(filepath)
	if err != nil {
		panic("Errow while opening file")
	}
	defer file.Close()
	filename := createBackupFileName(config)
	_, error := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:    aws.String(filename),
		Body:   file,
	})

	if error != nil {
		fmt.Println(error.Error())
		panic("Could not upload to S3.")
	}

}

// createBackupFileName creates the of the backup file.
func createBackupFileName(config config.Config) string{
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s/backup-%s.sql.gz", config.S3.DirectoryPrefix, strconv.FormatInt(timestamp, 16))
}