package cmd

import (
	"github.com/NiclasTimmeDev/pg-docker-backup/awsManager"
	"github.com/NiclasTimmeDev/pg-docker-backup/db"
	"github.com/NiclasTimmeDev/pg-docker-backup/fileManager"
	"github.com/spf13/cobra"
)

var backup = &cobra.Command{
	Use:   "backup",
	Short: "Create a backup.",
	Long: `Create a backup of your Postgres database that
		   lives in a Docker container and store it in S3.
		  `,
	Run: func(cmd *cobra.Command, args []string) {
		run(cmd)
	},
}

func init(){
	backup.PersistentFlags().String("container", "c", "The name of the container the postgres database lives in.")
	backup.PersistentFlags().String("username", "u", "The Postgres database username.")
	backup.PersistentFlags().String("database", "d", "The Postgres database name.")
}

func run(cmd *cobra.Command){
	containerName,_ := cmd.Flags().GetString("container")
	username,_ := cmd.Flags().GetString("username")
	dbName,_ := cmd.Flags().GetString("database")

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