package config

import "github.com/medtune/beta-platform/pkg/store/model"

type Meta struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	IsProd      bool   `yaml:"prod"`
}

type Server struct {
	Mode     string `yaml:"mode"`
	Port     int    `yaml:"port"`
	Protocol string `yaml:"protocol"`
	SSLTLS   bool   `yaml:"ssltls"`
}

type Database struct {
	Type    string `yaml:"type"`
	Prod    string `yaml:"prod"`
	Test    string `yaml:"test"`
	SSLMode int8   `yaml:"sslmode"`
	Creds   struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"creds"`
}

type Session struct {
	Type   string `yaml:"type"`
	Random bool   `yaml:"random"`
	Secret string `yaml:"secret"`
	Name   string `yaml:"name"`
}

type Crypto struct {
	Algo string `yaml:"algo"`
	Salt string `yaml:"salt"`
}

type PublicContent struct {
	Static string `yaml:"static`
}

type Secrets struct {
	Signup []string `yaml:"signup"`
}

type Create struct {
	Users []*model.User `yaml:"users"`
}

type StartupConfig struct {
	Meta     *Meta          `yaml:"meta"`
	Server   *Server        `yaml:"server"`
	Database *Database      `yaml:"database"`
	Session  *Session       `yaml:"session"`
	Crypto   *Crypto        `yaml:"crypto"`
	Public   *PublicContent `yaml:"public"`
	Secrets  *Secrets       `yaml:"secrets"`
	Create   *Create        `yaml:"create"`
}
