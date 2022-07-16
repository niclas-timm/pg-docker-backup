package fileManager

import (
	"compress/gzip"
	"io/ioutil"
	"os"
)


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