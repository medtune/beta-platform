// +build !prod !debug

package main

import (
	"runtime/pprof"

	"github.com/medtune/beta-platform/cmd/root"

	_ "github.com/medtune/beta-platform/cmd/copyright"
	_ "github.com/medtune/beta-platform/cmd/debug"
	_ "github.com/medtune/beta-platform/cmd/gen-config"
	_ "github.com/medtune/beta-platform/cmd/gen-views"
	_ "github.com/medtune/beta-platform/cmd/start"
	_ "github.com/medtune/beta-platform/cmd/sync-cxpba"
	_ "github.com/medtune/beta-platform/cmd/sync-db"
	_ "github.com/medtune/beta-platform/cmd/version"
)

func main() {

	defer pprof.StopCPUProfile()
	root.Execute()
}
