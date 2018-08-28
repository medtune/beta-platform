package config

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/store/model"
)

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
	Protocol string `yaml:"protocol,omitempty"`
	SSLTLS   bool   `yaml:"ssltls,omitempty"`
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
	MOC     int      `yaml:"max_idle_conns"`
	MIC     int      `yaml:"max_open_conns"`
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

type CustomModelConfig struct {
	ModelConfig
	PredictURI string
	HealthURI  string
}

// Capsul .
type Capsul struct {
	Inception *ModelConfig `yaml:"inception"`
	Mnist     *ModelConfig `yaml:"mnist"`
	Mura      *ModelConfig `yaml:"mura"`
	MuraIRNV2 *ModelConfig `yaml:"mura-irn-v2"`
	Chexray   *ModelConfig `yaml:"chexray"`
}

// CustomCapsul .
type CustomCapsul struct {
	Inception *ModelConfig `yaml:"inception"`
	Mnist     *ModelConfig `yaml:"mnist"`
	MuraCam   *ModelConfig `yaml:"mura-cam"`
	Mura      *ModelConfig `yaml:"mura"`
	Chexray   *ModelConfig `yaml:"chexray"`
}

// StartupConfig main configuration
type StartupConfig struct {
	Meta         *Meta          `yaml:"meta"`
	Server       *Server        `yaml:"server"`
	Database     *Database      `yaml:"database"`
	Session      *Session       `yaml:"session"`
	Crypto       *Crypto        `yaml:"crypto"`
	Public       *PublicContent `yaml:"public"`
	Secrets      *Secrets       `yaml:"secrets"`
	Create       *Create        `yaml:"create"`
	Capsul       *Capsul        `yaml:"capsul"`
	CustomCapsul *CustomCapsul  `yaml:"custom_capsul"`
}

// Validate configuration fields & subfields
func (sc *StartupConfig) Validate() (bool, error) {
	if sc.Meta == nil {
		return false, fmt.Errorf("Meta is nil")
	}

	if sc.Server == nil {
		return false, fmt.Errorf("Server is nil")
	}

	if sc.Database == nil {
		return false, fmt.Errorf("Database is nil")
	}

	if sc.Session == nil {
		return false, fmt.Errorf("Session is nil")
	}

	if sc.Crypto == nil {
		return false, fmt.Errorf("Crypto is nil")
	}

	if sc.Secrets == nil {
		return false, fmt.Errorf("Secrets is nil")
	}

	if sc.Capsul == nil {
		return false, fmt.Errorf("Capsul is nil")
	}

	if sc.CustomCapsul == nil {
		return false, fmt.Errorf("CustomCapsul is nil")
	}

	if sc.Create == nil {
		return false, fmt.Errorf("Create is nil")
	}

	return true, nil
}
