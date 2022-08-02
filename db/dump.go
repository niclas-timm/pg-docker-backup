package db

import (
	"fmt"
	"os/exec"

	"github.com/NiclasTimmeDev/pg-docker-backup/fileManager"
	"github.com/NiclasTimmeDev/pg-docker-backup/notifications"
)

// Dump creates a dump file from a docker container and stores it
// in the tmp directory.
func Dump(containerName string, username string, dbName string, table string) string {

	commandParams := make([]string, 0)
	commandParams = append(commandParams, "docker")
	commandParams = append(commandParams, "exec")
	commandParams = append(commandParams, containerName)
	commandParams = append(commandParams, "pg_dump")
	commandParams = append(commandParams, "-U")
	commandParams = append(commandParams, username)
	commandParams = append(commandParams, "--format=c")
	commandParams = append(commandParams, dbName)

	cmd := exec.Command("docker", "exec", containerName, "pg_dump", "-U", username, "--format=c", dbName)
	if table != "" {
		cmd = exec.Command("docker", "exec", containerName, "pg_dump", "-U", username, "--format=c", "--table", table, dbName)
	}

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

