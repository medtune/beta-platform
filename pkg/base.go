package pkg

import (
	"fmt"
	"strings"
)

var (
	Authors     string
	Owners      string
	LicenseURL  string
	LicenseType string

	VERSION     string
	LONGVERSION string
	Major       = "0"
	Minor       = "0"
	Patch       = "0"
	Revision    = "0-unknown"
	GitCommit   = "$Format:%H$"
	BuildDate   = "1970-01-01T00:00:00Z"
)

func init() {
	Owners = strings.Replace(strings.Replace(Owners, "/", ", ", -1), ".", " ", -1)
	Authors = strings.Replace(strings.Replace(Authors, "/", ", ", -1), ".", " ", -1)
	VERSION = fmt.Sprintf("v%s.%s.%s", Major, Minor, Patch)
	LONGVERSION = fmt.Sprintf("v%s.%s.%s-%s", Major, Minor, Patch, Revision)
}

// VersionInfo data
type VersionInfo struct {
	GitVersion string `json:"GitVersion"`
	GitCommit  string `json:"GitCommit"`
	Major      string `json:"Major"`
	Minor      string `json:"Minor"`
	Patch      string `json:"Patch"`
	Revision   string `json:"Revision"`
	Compiler   string `json:"Compiler"`
	GoVersion  string `json:"GoVersion"`
	Platform   string `json:"Platform"`
	BuildDate  string `json:"BuildDate"`
}

// CopyrightInfo data
type CopyrightInfo struct {
	Authors     string `json:"Authors"`
	Owners      string `json:"Owners"`
	LicenseType string `json:"LicenseType"`
	LicenseURL  string `json:"LicenseURL"`
}
