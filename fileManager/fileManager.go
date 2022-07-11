package fileManager

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

var TmpDirName = "tmp"


// CreateTmpDumpFile creates a gzipped file in the tmp directory
// from a slice of bytes.
func CreateTmpDumpFile(content []byte) error {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write(content)
	w.Close()

	// Create tmp directory if it not already exists.
	if _, err := os.Stat(TmpDirName); os.IsNotExist(err) {
		os.Mkdir(TmpDirName, os.ModePerm)
	}

	dumpFileName := fmt.Sprintf("%s/%s", TmpDirName, "temporary_dump.sql.gz")

	return ioutil.WriteFile(dumpFileName, b.Bytes(), 0666)
	
}