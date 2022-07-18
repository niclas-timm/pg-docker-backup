package cmd

import (
	"fmt"
	"os"

	"github.com/NiclasTimmeDev/pg-docker-backup/awsManager"
	"github.com/NiclasTimmeDev/pg-docker-backup/db"
	"github.com/NiclasTimmeDev/pg-docker-backup/fileManager"

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

// init method. Used for specifying the flags form the imp command.
func init(){
	imp.PersistentFlags().String("container", "c", "The name of the container the postgres database lives in.")
	imp.PersistentFlags().String("username", "u", "The name of the container the postgres database lives in.")
	imp.PersistentFlags().String("database", "d", "The name of the container the postgres database lives in.")
}

// runImport donwloads the latest dump file from S3 and dumps it intp the local database container.
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
	
	s3latestBackup := files[0]
	s3LatestBackupFilename := *s3latestBackup.Key

	// Create tmp directory if it not already exists.
	if _, err := os.Stat(fileManager.TmpImpDirName); os.IsNotExist(err) {
		os.Mkdir(fileManager.TmpImpDirName, os.ModePerm)
	}

	localTmpTargetFile, err := os.Create(fmt.Sprintf("%s/%s", fileManager.TmpImpDirName, fileManager.TmpImpFileName))

	defer localTmpTargetFile.Close()

	if err != nil {
		panic("Could not create temporary download file. Aborting.")
	}

	awsManager.DownloadS3Object(client, s3LatestBackupFilename, localTmpTargetFile)
	fileBytes, err := fileManager.UnzipFile(localTmpTargetFile)
	if err != nil {
		panic(err.Error())
	}

	os.WriteFile("tmp/import_file.sql", fileBytes, 0644)

	// Copy file into container with docker cp command
	containerName,_ := cmd.Flags().GetString("container")
	username,_ := cmd.Flags().GetString("username")
	dbName,_ := cmd.Flags().GetString("database")
	if containerName == "" || username == "" || dbName == "" {
		panic("Please provide containerName (--c), username (--u) and database name (--d)")
	}

	// Copy the import file into the container.
	copyImportFileOutput, err := fileManager.CopyFileToDockerContainer(containerName, "tmp/import_file.sql", "/home")
	if err != nil {
		fmt.Println(string(copyImportFileOutput))
		panic("Error while importing the database: Unable to copy the dump file into the target container.")
	}

	// Import the DB.
	importDumpOutput,_ := db.ImportDbDump(containerName, username, dbName)
	fmt.Println(string(importDumpOutput))
	
}
