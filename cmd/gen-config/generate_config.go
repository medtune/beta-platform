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

package genconfig

import (
	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/pkg/config"
	"github.com/spf13/cobra"
)

var (
	output string
)

func init() {
	genViewsCmd.Flags().StringVarP(&output, "output", "o", "gen.config.yml", "output directory")

	root.Cmd.AddCommand(genViewsCmd)
}

var genViewsCmd = &cobra.Command{
	Use:     "gen-views",
	Aliases: []string{"gen-config", "gen"},
	Short:   "Generate empty startup config file",
	Long:    `Generate empty startup config file`,
	Run: func(cmd *cobra.Command, args []string) {
		generateConfig()
	},
}

func generateConfig() {
	config.Generate(&config.StartupConfig{}, output)
}
