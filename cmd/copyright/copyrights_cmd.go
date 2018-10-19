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

package copyright

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/internal"
)

var (
	authors     bool
	owners      bool
	licenseType bool
	licenseURL  bool
)

func init() {
	copyrightCmd.Flags().BoolVarP(&owners, "owners", "o", false, "Print owners only")
	copyrightCmd.Flags().BoolVarP(&authors, "authors", "a", false, "Print authors only")
	copyrightCmd.Flags().BoolVarP(&licenseType, "license-type", "t", false, "Print license type")
	copyrightCmd.Flags().BoolVarP(&licenseURL, "license-url", "u", false, "Print license URL")
}

// versionCmd represents the version command
var copyrightCmd = &cobra.Command{
	Use:     "copyright",
	Aliases: []string{"cp", "license"},
	Short:   "Medtune beta-platform copyright",
	Long:    `Print Medtune Beta Platform Copyright informations`,
	Run: func(cmd *cobra.Command, args []string) {
		copyrightInfo := internal.GetCopyright()
		copyrightsJSON, _ := json.MarshalIndent(copyrightInfo, "", "    ")
		shouldNotPrintAll := authors || owners || licenseType || licenseURL
		if shouldNotPrintAll {
			if authors {
				fmt.Printf("%s\n", copyrightInfo.Authors)
			}
			if owners {
				fmt.Printf("%s\n", copyrightInfo.Owners)
			}
			if licenseType {
				fmt.Printf("%s\n", copyrightInfo.LicenseType)
			}
			if licenseURL {
				fmt.Printf("%s\n", copyrightInfo.LicenseURL)
			}
		} else {
			fmt.Printf("%s\n", string(copyrightsJSON))
		}
	},
}

func init() {
	root.Cmd.AddCommand(copyrightCmd)
}
