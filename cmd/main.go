package main

import (
	_ "github.com/medtune/beta-platform/cmd/capsules"
	_ "github.com/medtune/beta-platform/cmd/debug"
	_ "github.com/medtune/beta-platform/cmd/gen-views"
	_ "github.com/medtune/beta-platform/cmd/start"
	_ "github.com/medtune/beta-platform/cmd/syncdb"
	_ "github.com/medtune/beta-platform/cmd/version"

	"github.com/medtune/beta-platform/cmd/root"
)

func main() {
	root.Execute()
}
