package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/NiclasTimmeDev/pg-docker-backup/utils"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

// ParseConfigFiles parses multiple yml files and merges them into one struct.
func ParseYmlFiles[T any](bindObj *T, filenames ...string) error {
	var resultValues map[string]interface{}
    for _, filename := range filenames {

		if _, err := os.Stat(TmpDirName); os.IsNotExist(err) {
			continue
		}

        var override map[string]interface{}
        bs, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Print(err.Error())
            continue
        }
        if err := yaml.Unmarshal(bs, &override); err != nil {
            fmt.Print(err.Error())
            continue
        }

        //check if is nil. This will only happen for the first filename
        if resultValues == nil {
            resultValues = override
        } else {
            resultValues = utils.DeepMergeMaps(resultValues, override)
        }

    }

    mapstructure.Decode(resultValues, &bindObj);

    return nil
}