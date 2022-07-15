package cmd

import (
	"fmt"
	"pg-docker-backup/awsManager"
	"pg-docker-backup/db"
	"pg-docker-backup/fileManager"

	"github.com/spf13/cobra"
)

var backup = &cobra.Command{
	Use:   "backup",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd)
	},
}

func init(){
	backup.PersistentFlags().String("container", "c", "The name of the container the postgres database lives in.")
	backup.PersistentFlags().String("username", "u", "The name of the container the postgres database lives in.")
	backup.PersistentFlags().String("database", "d", "The name of the container the postgres database lives in.")
}

func run(cmd *cobra.Command){
	containerName,_ := cmd.Flags().GetString("container")
	username,_ := cmd.Flags().GetString("username")
	dbName,_ := cmd.Flags().GetString("database")
	fmt.Print(containerName)

	if containerName == "" || username == "" || dbName == "" {
		panic("Please provide containerName (--c), username (--u) and database name (--d)")
	}

	// Create the database dump from the docker container based on the
	// inputs the user gave when executing the script.
	tmpFilename := db.Dump(containerName, username, dbName)

	// Connect to S3 and store the dump file there.
	// Plus, delete excess backup files from S3 in order to save storage space.
	client := awsManager.ConnectToS3()
	existingBucketItems := awsManager.GetAllBackupsFromS3(client)
	awsManager.DeleteExcessBackupsFromS3(client, existingBucketItems)
	awsManager.UploadToS3(client, tmpFilename)

	// Delet the temporary dump file since we don't need it anymore
	// after uploading it to S3.
	fileManager.RemoveTmpDumpFile(tmpFilename)
}