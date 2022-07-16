package db

import "os/exec"


func ImportDbDump(containerName string, username string, dbName string) ([]byte, error) {
	shellCommand := exec.Command(
		"docker",
		"exec",
		containerName,
		"pg_restore",
		"--verbose",
		"--clean",
		"--no-acl",
		"--no-owner",
		"-U",
		username,
		"-d",
		dbName,
		"/home/import_file.dump",
	)

	return shellCommand.CombinedOutput()
}