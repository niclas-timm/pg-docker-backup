package main

import (
	"flag"
	"pg-docker-backup/awsManager"
	"pg-docker-backup/db"
	"pg-docker-backup/fileManager"

	"github.com/joho/godotenv"
)

// main is the entry method of the program.
func main(){

	godotenv.Load()

	// Capture relevant data from command line arguments
	var containerName string
	var username string
	var dbName string
	flag.StringVar(&containerName, "c", "", "The name of the container the postgres database lives in.")
	flag.StringVar(&username, "u", "", "The postgres user name.")
	flag.StringVar(&dbName, "d", "", "The postgres database name.")
	flag.Parse()

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