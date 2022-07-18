package db

import (
	"fmt"
	"os/exec"

	"github.com/NiclasTimmeDev/pg-docker-backup/fileManager"
	"github.com/NiclasTimmeDev/pg-docker-backup/notifications"
)


var tmpDirName = "tmp"

// Dump creates a dump file from a docker container and stores it
// in the tmp directory.
func Dump(containerName string, username string, dbName string) string{
	cmd := exec.Command("docker", "exec", containerName, "pg_dump", "-U", username, "--format=c", dbName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		notifications.NotifyViaAllChannels(string(output))
		panic("Error while creating dump")
	}
	
	filename, e := fileManager.CreateTmpDumpFile(output)
	if e != nil {
		fmt.Println(e.Error())
		notifications.NotifyViaAllChannels(e.Error())
		panic("Aborting.")
	}

	return filename
}

