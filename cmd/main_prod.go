// +build prod

package main

import (
	_ "github.com/medtune/beta-platform/cmd/start"
	_ "github.com/medtune/beta-platform/cmd/version"

	"github.com/medtune/beta-platform/cmd/root"
)

func main() {
	root.Execute()
}
