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

package debug

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/server"
)

var (
	static string
	port   int
)

func init() {
	debugCmd.Flags().StringVarP(&static, "static", "s", "./static", "Static files directory")
	debugCmd.Flags().IntVarP(&port, "port", "p", 8005, "port")

	root.Cmd.AddCommand(debugCmd)
}

// startCmd represents the start command
var debugCmd = &cobra.Command{
	Use:     "debug",
	Aliases: []string{"debug-server"},
	Short:   "debug server for UI dev",
	Long:    `Debug UI server`,
	Run: func(cmd *cobra.Command, args []string) {
		debugServer()
	},
}

func debugServer() {
	gin.SetMode(gin.DebugMode)

	Server := server.Debug(
		static,
		port,
	)

	if err := Server.Run(); err != nil {
		log.Fatal(err)
	}
}
