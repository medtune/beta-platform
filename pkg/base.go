package pkg

import (
	"fmt"
	"strings"
)

var (
	Authors     string = ""
	Owners      string = ""
	LicenseURL  string = ""
	LicenseType string = ""

	VERSION     string = ""
	LONGVERSION string = ""
	Major       string = "0"
	Minor       string = "0"
	Patch       string = "0"
	Revision    string = "0-unknown"
	GitCommit   string = "$Format:%H$"
	BuildDate   string = "1970-01-01T00:00:00Z"
)

func init() {
	Owners = strings.Replace(Owners, ".", " ", -1)
	Authors = strings.Replace(Authors, ".", " ", -1)
	VERSION = fmt.Sprintf("v%s.%s.%s", Major, Minor, Patch)
	LONGVERSION = fmt.Sprintf("v%s.%s.%s+%s", Major, Minor, Patch, Revision)
}

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

type CopyrightInfo struct {
	Authors     string `json:"Authors`
	Owners      string `json:"Owners"`
	LicenseType string `json:"LicenseType"`
	LicenseURL  string `json:"LicenseURL"`
}
