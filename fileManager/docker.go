package fileManager

import (
	"fmt"
	"os/exec"
)

// CopyFileToDockerContainer copies a local file into a docker container.
func CopyFileToDockerContainer(containerName string, srcFile string, targetPath string) ([]byte, error) {
	copyImportFile := exec.Command(
		"docker",
		"cp",
		srcFile,
		fmt.Sprintf("%s:%s", containerName, targetPath),
	)

	return copyImportFile.CombinedOutput()
}