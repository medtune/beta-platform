package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	configPath    = "config.yml"
	configPathGen = "config-gen.yml"
)

// SetDefaultPaths .
func SetDefaultPaths(path string) {
	configPath = path + ".yml"
	configPathGen = path + "-gen.yml"
}

// LoadConfigDefault .
func LoadConfigDefault() (*StartupConfig, error) {
	return LoadConfigFromPath(configPath)
}

// LoadConfigFromPath .
func LoadConfigFromPath(path string) (*StartupConfig, error) {
	var config StartupConfig
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
