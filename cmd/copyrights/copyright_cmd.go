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

package copyrights

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/pkg"
)

// versionCmd represents the version command
var copyrightsCmd = &cobra.Command{
	Use:     "copyrights",
	Aliases: []string{"cprs", "license"},
	Short:   "Medtune beta-platform copyrights",
	Long:    `Print Medtune Beta Platform Copyrights informations`,
	Run: func(cmd *cobra.Command, args []string) {
		copyrightsJSON, _ := json.MarshalIndent(pkg.GetCopyrights(), "", "    ")
		fmt.Printf("%s\n", copyrightsJSON)
	},
}

func init() {
	root.Cmd.AddCommand(copyrightsCmd)
}
