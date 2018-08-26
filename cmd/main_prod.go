// +build prod

package main

import (
	"github.com/medtune/beta-platform/cmd/root"

	_ "github.com/medtune/beta-platform/cmd/start"
	_ "github.com/medtune/beta-platform/cmd/version"
)

func main() {
	root.Execute()
}
