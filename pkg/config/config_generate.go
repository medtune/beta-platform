package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// GenDefault generate empty config file with default name
func GenDefault() error {
	return Generate(&StartupConfig{}, configPathGen)
}

// Generate empty config file
func Generate(config *StartupConfig, filename string) error {
	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, yamlBytes, 0644)
	return err
}
