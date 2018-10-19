package config

import (
	"fmt"

	"github.com/medtune/beta-platform/internal/store/model"
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
	MOC     int      `yaml:"max_idle_conns"`
	MIC     int      `yaml:"max_open_conns"`
	LTC     float64  `yaml:"max_conn_lifetime"`
	Creds   *DBCreds `yaml:"creds"`
}

// Session .
type Session struct {
	Type     string `yaml:"type"`
	Random   bool   `yaml:"random"`
	Secret   string `yaml:"secret"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Address  string `yaml:"address"`
}

// Crypto .
type Crypto struct {
	Algo string `yaml:"algo"`
	Salt string `yaml:"salt"`
}

// PublicContent .
type PublicContent struct {
	Static string `yaml:"static"`
	Prefix string `yaml:"prefix"`
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
	Name      string `yaml:"name"`
	Model     string `yaml:"model"`
	Signature string `yaml:"signature"`
	Version   int    `yaml:"version"`
	Address   string `yaml:"address"`
}

// Capsul .
type Capsul struct {
	Inception    *ModelConfig `yaml:"inception"`
	Mnist        *ModelConfig `yaml:"mnist"`
	MuraMNV2     *ModelConfig `yaml:"mura-mn-v2"`
	MuraIRNV2    *ModelConfig `yaml:"mura-irn-v2"`
	ChexrayMNV2  *ModelConfig `yaml:"chexray-mn-v2"`
	ChexrayDN121 *ModelConfig `yaml:"chexray-dn-121"`
}

// CustomCapsul .
type CustomCapsul struct {
	MuraMNV2Cam     *ModelConfig `yaml:"mura-mn-v2-cam"`
	MuraIRNV2Cam    *ModelConfig `yaml:"mura-irn-v2-cam"`
	ChexrayMNV2Cam  *ModelConfig `yaml:"chexray-mn-v2-cam"`
	ChexrayDN121Cam *ModelConfig `yaml:"chexray-dn-121-cam"`
	ChexrayPP       *ModelConfig `yaml:"chexray-pp-helper"`
}

// StartupConfig main configuration
type StartupConfig struct {
	Meta         *Meta          `yaml:"meta,omitempty"`
	Server       *Server        `yaml:"server,omitempty"`
	Database     *Database      `yaml:"database,omitempty"`
	Session      *Session       `yaml:"session,omitempty"`
	Crypto       *Crypto        `yaml:"crypto,omitempty"`
	Public       *PublicContent `yaml:"public,omitempty"`
	Secrets      *Secrets       `yaml:"secrets,omitempty"`
	Create       *Create        `yaml:"create,omitempty"`
	Capsul       *Capsul        `yaml:"capsul,omitempty"`
	CustomCapsul *CustomCapsul  `yaml:"custom_capsul,omitempty"`
}

// Validate configuration fields & subfields
func (sc *StartupConfig) Validate() (bool, error) {
	if sc.Meta == nil {
		return false, fmt.Errorf("meta object is nil")
	}

	if sc.Server == nil {
		//return false, fmt.Errorf("Server is nil")
	}

	if sc.Database == nil {
		return false, fmt.Errorf("database object is nil")
	}

	if sc.Session == nil {
		return false, fmt.Errorf("session object is nil")
	}

	if sc.Crypto == nil {
		return false, fmt.Errorf("crypto object is nil")
	}

	if sc.Secrets == nil {
		return false, fmt.Errorf("secrets object is nil")
	}

	if sc.Capsul == nil {
		return false, fmt.Errorf("capsul object is nil")
	}

	if sc.CustomCapsul == nil {
		return false, fmt.Errorf("custom capsul object is nil")
	}

	if sc.Create == nil {
		return false, fmt.Errorf("create object is nil")
	}

	return true, nil
}
