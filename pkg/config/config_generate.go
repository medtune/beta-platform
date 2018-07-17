package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func GenDefault() error {
	return Gen(StartupConfig{}, configPathGen)
}

func Gen(config StartupConfig, filename string) error {
	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, yamlBytes, 0644)
	return err
}
