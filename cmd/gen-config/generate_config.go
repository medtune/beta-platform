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
	"log"
	"os"

	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/pkg"
	"github.com/medtune/beta-platform/pkg/config"
	"github.com/medtune/beta-platform/pkg/store/model"
	"github.com/spf13/cobra"
)

var (
	output string
	empty  bool
)

func init() {
	genConfigCmd.Flags().StringVarP(&output, "output", "o", "gen.config.yml", "output directory")
	genConfigCmd.Flags().BoolVarP(&empty, "empty", "e", false, "empty configuration")

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
	if empty {
		err := config.GenDefault()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	err := config.Generate(&config.StartupConfig{
		Meta: &config.Meta{
			Name:    "Medtune beta platform",
			Version: pkg.VERSION,
			IsProd:  true,
		},
		Server: &config.Server{},
		Database: &config.Database{
			Type: "postgres",
			Prod: "medtune",
			Test: "medtune-test",
			Creds: &config.DBCreds{
				Host:     "localhost",
				Port:     5432,
				User:     "mdtn",
				Password: "mdtn",
			},
			SSLMode: 0,
			MOC:     100,
			MIC:     100,
		},
		Session: &config.Session{
			Type:   "cookie",
			Random: true,
		},
		Crypto: &config.Crypto{
			Algo: "SHA256",
			Salt: "DEFAULT-SALT",
		},
		Public: &config.PublicContent{
			Static: "./static",
		},
		Secrets: &config.Secrets{
			Signup: []string{
				"supersecret",
				"siistrasbourg",
			},
		},
		Create: &config.Create{
			Users: []*model.User{
				{
					Username: "admin",
					Email:    "admin@medtune.eu",
					Password: "admin",
				},
				{
					Username: "test",
					Email:    "test@medtune.eu",
					Password: "test",
				},
			},
		},
		Capsul: &config.Capsul{
			Mnist: &config.ModelConfig{
				Model:     "mnist",
				Signature: "predict_images",
				Version:   1,
				Address:   "localhost:9000",
			},
			Inception: &config.ModelConfig{
				Model:     "inception",
				Signature: "predict_images",
				Version:   1,
				Address:   "localhost:9001",
			},
			Mura: &config.ModelConfig{
				Model:     "mura",
				Signature: "predict_images",
				Version:   1,
				Address:   "localhost:9002",
			},
			Chexray: &config.ModelConfig{
				Model:     "chexray",
				Signature: "predict_images",
				Version:   1,
				Address:   "localhost:9003",
			},
		},
	}, output)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
