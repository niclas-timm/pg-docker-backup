package fileManager

import (
	"compress/gzip"
	"io/ioutil"
	"os"
)

// UnzipFile unzips a gzipped file and returns its content as an array of bytes.
func UnzipFile(file *os.File) ([]byte, error){
	defer file.Close()

	fz, err := gzip.NewReader(file)
    if err != nil {
        return nil, err
    }
    defer fz.Close()
	
	s, err := ioutil.ReadAll(fz)
    if err != nil {
        return nil, err
    }
    return s, nil   
}