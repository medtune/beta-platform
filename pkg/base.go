package pkg

import "fmt"

var (
	VERSION   string
	Major     string = "0"
	Minor     string = "0"
	Patch     string = "0-unknown"
	GitCommit string = "$Format:%H$"
	BuildDate string = "1970-01-01T00:00:00Z"
)

func init() {
	VERSION = fmt.Sprintf("v%s.%s.%s", Major, Minor, Patch)
}

type VersionInfo struct {
	Major      string `json:"Major"`
	Minor      string `json:"Minor"`
	Patch      string `json:"Patch"`
	GitVersion string `json:"GitVersion"`
	GitCommit  string `json:"GitCommit"`
	GoVersion  string `json:"GoVersion"`
	Platform   string `json:"Platform"`
	Compiler   string `json:"Compiler"`
	BuildDate  string `json:"BuildDate"`
}
