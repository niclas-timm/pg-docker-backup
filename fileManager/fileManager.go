package fileManager

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/NiclasTimmeDev/pg-docker-backup/config"
	"github.com/NiclasTimmeDev/pg-docker-backup/notifications"
)

// CreateTmpDumpFile creates a gzipped file in the tmp directory
// from a slice of bytes.
func CreateTmpDumpFile(content []byte) (string, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(content)
	w.Close()

	// Create tmp directory if it not already exists.
	if _, err := os.Stat(config.TmpDirName); os.IsNotExist(err) {
		os.Mkdir(config.TmpDirName, os.ModePerm)
	}

	filename := createTmpBackupFileName()

	return filename, ioutil.WriteFile(filename, b.Bytes(), 0666)
	
}

// RemoveTmpDumpFile removes the temporary dump file.
// 
// The dump file is created so that it can be uploaded to 
// S3. Afterwards, we don't need it anymore so we can delete it.
func RemoveTmpDumpFile(filename string){
	err := os.Remove(filename)
	if err != nil {
		errMsg := fmt.Sprintf("Could not remove temporary dump file: %s", filename)
		notifications.NotifyViaAllChannels(errMsg)
		panic(errMsg)
	}
}

func CreateLocalTmpImportDump(){}

// createTmpBackupFileName creates the name for the temporary backup file.
func createTmpBackupFileName() string{
	timestamp := time.Now().Unix()
	return fmt.Sprintf("%s/backup-%s.sql.gz", config.TmpDirName, strconv.FormatInt(timestamp, 16))
}

