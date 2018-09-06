// +build debug

package main

import (
	"github.com/medtune/beta-platform/cmd/root"

	_ "github.com/medtune/beta-platform/cmd/copyright"
	_ "github.com/medtune/beta-platform/cmd/debug"
	_ "github.com/medtune/beta-platform/cmd/version"
)

func main() {
	root.Execute()
}
