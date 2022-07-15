package awsManager

import (
	"context"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// GetAllBackupsFromS3 returns all backup files from an the S3 bucket.
//
// The items will be returned in descending order by modification date.
// This means, "newer" items will come first in the slice.
func GetAllBackupsFromS3(c *s3.Client) []types.Object{
	output, err := c.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Prefix: aws.String(BackupDirectory),
	})

	if err != nil {
		panic("Unable fo fetch bucket items.")
	}

	SortS3ItemsByTime(output.Contents)	

	return output.Contents
}

// SortS3ItemsByTime sorts a slice of S3 items
//
// The method will sort the items descending by last modification date.
// This means, "newer" items will come first in the slice.
func  SortS3ItemsByTime(items []types.Object){
	sort.SliceStable(items, func(i, j int) bool {
		currentTime := *items[i].LastModified
		nextTime := *items[j].LastModified

		return currentTime.After(nextTime)
	})
}