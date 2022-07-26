package db

import "os/exec"

// ImportDbDump runs a shell command that imports a database dump into a postgres
// database that lives inside a docker container.
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
		"/home/import_file.sql",
	)

	return shellCommand.CombinedOutput()
}