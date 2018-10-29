package internal

import (
	"fmt"
	"runtime"
)

var (
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

// GetVersion .
func GetVersion() *VersionInfo {
	return &VersionInfo{
		Major:      Major,
		Minor:      Minor,
		Patch:      Patch,
		Revision:   Revision,
		GitVersion: LONGVERSION,
		GitCommit:  GitCommit,
		BuildDate:  BuildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
