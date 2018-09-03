package config

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/store/model"
)

// Meta .
type Meta struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Version     string `yaml:"version,omitempty"`
	IsProd      bool   `yaml:"prod,omitempty"`
}

// Server .
type Server struct {
	Mode     string `yaml:"mode,omitempty"`
	Port     int    `yaml:"port,omitempty"`
	Protocol string `yaml:"protocol,omitempty"`
	SSLTLS   bool   `yaml:"ssltls,omitempty"`
}

// DBCreds .
type DBCreds struct {
	Host     string `yaml:"host,omitempty"`
	Port     int    `yaml:"port,omitempty"`
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
}

// Database .
type Database struct {
	Type    string   `yaml:"type,omitempty"`
	Prod    string   `yaml:"prod,omitempty"`
	Test    string   `yaml:"test,omitempty"`
	SSLMode int8     `yaml:"sslmode,omitempty"`
	MOC     int      `yaml:"max_idle_conns,omitempty"`
	MIC     int      `yaml:"max_open_conns,omitempty"`
	Creds   *DBCreds `yaml:"creds,omitempty"`
}

// Session .
type Session struct {
	Type     string `yaml:"type,omitempty"`
	Random   bool   `yaml:"random,omitempty"`
	Secret   string `yaml:"secret,omitempty"`
	Name     string `yaml:"name,omitempty"`
	Password string `yaml:"password,omitempty"`
	Database string `yaml:"database,omitempty"`
	Address  string `yaml:"address,omitempty"`
}

// Crypto .
type Crypto struct {
	Algo string `yaml:"algo,omitempty"`
	Salt string `yaml:"salt,omitempty"`
}

// PublicContent .
type PublicContent struct {
	Static string `yaml:"static,omitempty"`
}

// Secrets .
type Secrets struct {
	Signup []string `yaml:"signup,omitempty"`
}

// Create .
type Create struct {
	Users []*model.User `yaml:"users,omitempty"`
}

// ModelConfig .
type ModelConfig struct {
	Model     string `yaml:"model,omitempty"`
	Signature string `yaml:"signature,omitempty"`
	Version   int    `yaml:"version,omitempty"`
	Address   string `yaml:"address,omitempty"`
}

// Capsul .
type Capsul struct {
	Inception    *ModelConfig `yaml:"inception,omitempty"`
	Mnist        *ModelConfig `yaml:"mnist,omitempty"`
	MuraMNV2     *ModelConfig `yaml:"mura-mn-v2,omitempty"`
	MuraIRNV2    *ModelConfig `yaml:"mura-irn-v2,omitempty"`
	ChexrayMNV2  *ModelConfig `yaml:"chexray-mn-v2,omitempty"`
	ChexrayDN121 *ModelConfig `yaml:"chexray-dn-121,omitempty"`
}

// CustomCapsul .
type CustomCapsul struct {
	MuraMNV2Cam     *ModelConfig `yaml:"mura-mn-v2-cam,omitempty"`
	MuraIRNV2Cam    *ModelConfig `yaml:"mura-irn-v2-cam,omitempty"`
	ChexrayMNV2Cam  *ModelConfig `yaml:"chexray-mn-v2-cam,omitempty"`
	ChexrayDN121Cam *ModelConfig `yaml:"chexray-dn-121-cam,omitempty"`
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
		return false, fmt.Errorf("Meta is nil")
	}

	/*
		if sc.Server == nil {
			return false, fmt.Errorf("Server is nil")
		}
	*/

	if sc.Database == nil {
		return false, fmt.Errorf("database is nil")
	}

	if sc.Session == nil {
		return false, fmt.Errorf("session is nil")
	}

	if sc.Crypto == nil {
		return false, fmt.Errorf("crypto is nil")
	}

	if sc.Secrets == nil {
		return false, fmt.Errorf("secrets is nil")
	}

	if sc.Capsul == nil {
		return false, fmt.Errorf("capsul is nil")
	}

	if sc.CustomCapsul == nil {
		return false, fmt.Errorf("custom capsul is nil")
	}

	if sc.Create == nil {
		return false, fmt.Errorf("create is nil")
	}

	return true, nil
}
