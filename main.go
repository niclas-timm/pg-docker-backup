package main

import (
	"flag"
	"pg-docker-backup/awsManager"
	"pg-docker-backup/db"

	"github.com/joho/godotenv"
)

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

	db.Dump(containerName, username, dbName)
	client := awsManager.ConnectToS3()
	awsManager.UploadToS3(client, "tmp/temporary_dump.sql.gz")
}