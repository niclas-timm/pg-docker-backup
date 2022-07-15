package cmd

import (
	"context"
	"fmt"
	"os"
	"pg-docker-backup/awsManager"
	"pg-docker-backup/fileManager"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cobra"
)

var imp = &cobra.Command{
	Use:   "import",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		runImport(cmd)
	},
}

func init(){
	// TODO: Set flags.
}

func runImport(cmd *cobra.Command){
	client := awsManager.ConnectToS3()
	files := awsManager.GetAllBackupsFromS3(client)

	if len(files) == 0 {
		panic("No backups stored in S3. Aborting.")
	}

	
	// Create tmp directory if it not already exists.
	if _, err := os.Stat(fileManager.TmpImpDirName); os.IsNotExist(err) {
		os.Mkdir(fileManager.TmpImpDirName, os.ModePerm)
	}
	
	latestBackup := files[0]

	localFileName := *latestBackup.Key

	// Create tmp directory if it not already exists.
	if _, err := os.Stat(fileManager.TmpImpDirName); os.IsNotExist(err) {
		os.Mkdir(fileManager.TmpImpDirName, os.ModePerm)
	}

	localFile, err := os.Create(fmt.Sprintf("%s/%s", fileManager.TmpImpDirName, fileManager.TmpImpFileName))

	defer localFile.Close()

	if err != nil {
		panic("Could not create temporary download file. Aborting.")
	}

	downloader := manager.NewDownloader(client)

    _, err = downloader.Download(context.TODO(), localFile, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key: aws.String(localFileName),
	})

	if err != nil {
		panic(err.Error())
	}

	fmt.Print("Successfully downloaded backup from S3.")

}
