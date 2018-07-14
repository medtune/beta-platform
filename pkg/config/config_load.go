package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	configPath    = "config.yaml"
	configPathGen = "config-gen.yaml"
)

func SetDefaultPaths(path string) {
	configPath = path + ".yaml"
	configPathGen = path + "-gen.yaml"
}

func LoadConfigDefault() (*StartupConfig, error) {
	return LoadConfigFromPath(configPath)
}

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
