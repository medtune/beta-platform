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
	genConfigCmd.Flags().StringVarP(&output, "output", "o", "gen.config.yml", "output directory")

	root.Cmd.AddCommand(genConfigCmd)
}

var genConfigCmd = &cobra.Command{
	Use:     "gen-config",
	Aliases: []string{"gen-cfg"},
	Short:   "Generate empty startup config file",
	Long:    `Generate empty startup config file`,
	Run: func(cmd *cobra.Command, args []string) {
		generateConfig()
	},
}

func generateConfig() {
	config.Generate(&config.StartupConfig{
		Meta:   &config.Meta{},
		Server: &config.Server{},
		Database: &config.Database{
			Creds: &config.DBCreds{},
		},
		Session: &config.Session{},
		Crypto:  &config.Crypto{},
		Public:  &config.PublicContent{},
		Secrets: &config.Secrets{},
		Create:  &config.Create{},
		Capsul: &config.Capsul{
			Inception: &config.ModelConfig{},
			Mnist:     &config.ModelConfig{},
			Mura:      &config.ModelConfig{},
			Chexray:   &config.ModelConfig{},
		},
	}, output)
}
