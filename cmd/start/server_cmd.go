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

package start

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/pkg/config"
	"github.com/medtune/beta-platform/pkg/initpkg"
	"github.com/medtune/beta-platform/server"
	"github.com/spf13/cobra"
)

var (
	static     string
	configFile string
	port       int
	ginMode    int
)

func init() {
	startCmd.Flags().StringVarP(&configFile, "file", "f", "config.yaml", "Configuration file name")
	startCmd.Flags().StringVarP(&static, "static", "s", "./static", "Static files directory")

	startCmd.Flags().IntVarP(&port, "port", "p", 8005, "port")
	startCmd.Flags().IntVarP(&ginMode, "gin-mode", "g", 0, "Gin server mode")

	root.Cmd.AddCommand(startCmd)
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run", "run-server"},
	Short:   "Run Medtune beta server",
	Long:    `Run Medtune beta server`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func runServer() {
	if ginMode == 0 {
		gin.SetMode(gin.DebugMode)
	} else if ginMode == 1 {
		gin.SetMode(gin.ReleaseMode)
	}

	var configuration *config.StartupConfig

	if cfg, err := config.LoadConfigFromPath(configFile); err != nil {
		log.Panic(err)
	} else {
		configuration = cfg
	}

	if err := initpkg.InitFromConfig(configuration); err != nil {
		log.Panic(err)
	}

	Server := server.New(
		static,
		port,
	)

	Server.Run()
}
