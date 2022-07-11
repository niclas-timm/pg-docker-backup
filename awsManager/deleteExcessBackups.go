package awsManager

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// DeleteExcessBackupsFromS3 deletes excess backup files from S3.
//
// To keep the storage size as small as possible, we allow a maximum
// of 7 backup files in S3. As we also want to upload one new file,
// this method deletes all backup files from S3 but the newest 6 ones
// (because 6 + the new one that will be created = 7).
func DeleteExcessBackupsFromS3(client *s3.Client, items []types.Object){
	if len(items) < maxNumberOfS3Backups {
		return
	}

	numberToKeep := maxNumberOfS3Backups - 1;
	remove := items[numberToKeep:]
	for _, itemToRemove := range remove {
		_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
			Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
			Key: itemToRemove.Key,
		})

		if err != nil {
			fmt.Print(err.Error())
		}
	}
}