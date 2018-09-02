package pkg

import (
	"fmt"
	"os"
	"runtime"
)

// GetVersion .
func GetVersion() *VersionInfo {
	return &VersionInfo{
		Major:      Major,
		Minor:      Minor,
		Patch:      Patch,
		GitVersion: VERSION,
		GitCommit:  GitCommit,
		BuildDate:  BuildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// PrintAndExit will print the version and exit.
func PrintAndExit() {
	fmt.Println(VERSION)
	os.Exit(0)
}
