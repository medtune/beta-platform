package config

import "github.com/medtune/beta-platform/pkg/store/model"

// Meta .
type Meta struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	IsProd      bool   `yaml:"prod"`
}

// Server .
type Server struct {
	Mode     string `yaml:"mode"`
	Port     int    `yaml:"port"`
	Protocol string `yaml:"protocol"`
	SSLTLS   bool   `yaml:"ssltls"`
}

// DBCreds .
type DBCreds struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// Database .
type Database struct {
	Type    string   `yaml:"type"`
	Prod    string   `yaml:"prod"`
	Test    string   `yaml:"test"`
	SSLMode int8     `yaml:"sslmode"`
	Creds   *DBCreds `yaml:"creds"`
}

// Session .
type Session struct {
	Type   string `yaml:"type"`
	Random bool   `yaml:"random"`
	Secret string `yaml:"secret"`
	Name   string `yaml:"name"`
}

// Crypto .
type Crypto struct {
	Algo string `yaml:"algo"`
	Salt string `yaml:"salt"`
}

// PublicContent .
type PublicContent struct {
	Static string `yaml:"static"`
}

// Secrets .
type Secrets struct {
	Signup []string `yaml:"signup"`
}

// Create .
type Create struct {
	Users []*model.User `yaml:"users"`
}

// ModelConfig .
type ModelConfig struct {
	Model     string `yaml:"model"`
	Signature string `yaml:"signature"`
	Version   int    `yaml:"version"`
	Address   string `yaml:"address"`
}

// Capsul .
type Capsul struct {
	Inception *ModelConfig `yaml:"inception"`
	Mnist     *ModelConfig `yaml:"mnist"`
	Mura      *ModelConfig `yaml:"mura"`
	Chexray   *ModelConfig `yaml:"chexray"`
}

// StartupConfig main configuration
type StartupConfig struct {
	Meta     *Meta          `yaml:"meta"`
	Server   *Server        `yaml:"server"`
	Database *Database      `yaml:"database"`
	Session  *Session       `yaml:"session"`
	Crypto   *Crypto        `yaml:"crypto"`
	Public   *PublicContent `yaml:"public"`
	Secrets  *Secrets       `yaml:"secrets"`
	Create   *Create        `yaml:"create"`
	Capsul   *Capsul        `yaml:"capsul"`
}
