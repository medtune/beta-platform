package data

import "github.com/medtune/beta-platform/internal/config"

// Dashboard .
type Dashboard struct {
	Status      string
	Title       string
	Config      string
	Version     string
	StartupDate string
	Capsules    []*config.ModelConfig
	Storage     []*StorageConfig
}

// StorageConfig .
type StorageConfig struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
