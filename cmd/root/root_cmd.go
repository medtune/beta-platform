// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "medtune-beta",
	Short: "Medtune Beta Platform entrypoint",
	Long:  `Medtune Platform ~ Beta`,
}

// Execute root cmd
func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Printf("[FATAL] %v", err)
		os.Exit(1)
	}
}
