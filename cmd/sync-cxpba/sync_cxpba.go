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

package synccxpba

import (
	"log"

	"github.com/medtune/beta-platform/hack/xlsx2pg/xlsx2pg"
	"github.com/spf13/cobra"

	"github.com/medtune/beta-platform/cmd/root"
)

var (
	configFile string
	CXPBAFile  string
)

func init() {
	syncCXPBAcmd.Flags().StringVarP(&configFile, "config", "c", "config.yml", "Configuration file name")
	syncCXPBAcmd.Flags().StringVarP(&CXPBAFile, "xlsx", "x", "CXPBA.xlsx", "CXPBA excel file")

	root.Cmd.AddCommand(syncCXPBAcmd)
}

// syncdbCmd represents the syncdb command
var syncCXPBAcmd = &cobra.Command{
	Use:     "sync-cxpba",
	Aliases: []string{"synccxpba"},
	Short:   "Automigrate & sync CXPBA data",
	Run: func(cmd *cobra.Command, args []string) {
		if err := xlsx2pg.SyncCXPBAexcel(CXPBAFile, configFile); err != nil {
			log.Fatal(err)
		}
	},
}
