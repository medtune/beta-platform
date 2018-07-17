package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	configPath    = "config.yml"
	configPathGen = "config-gen.yml"
)

func SetDefaultPaths(path string) {
	configPath = path + ".yml"
	configPathGen = path + "-gen.yml"
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
