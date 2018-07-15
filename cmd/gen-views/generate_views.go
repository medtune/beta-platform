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

package genviews

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/pkg/tmpl"
	"github.com/medtune/beta-platform/pkg/tmpl/data"
	"github.com/spf13/cobra"
)

var (
	output string
	views  string
)

var (
	TEMPLATES map[string]*template.Template
)

func init() {
	TEMPLATES = tmpl.GetTemplatesMap()

	genViewsCmd.Flags().StringVarP(&output, "output", "o", "generated-views", "output directory (default generate-views)")
	genViewsCmd.Flags().StringVarP(&views, "views", "v", "...", "views to generate (comma separated string)")

	root.Cmd.AddCommand(genViewsCmd)
}

// startCmd represents the start command
var genViewsCmd = &cobra.Command{
	Use:     "gen-views",
	Aliases: []string{"gen-tmpl", "gen"},
	Short:   "generate views html files",
	Long:    `generate views html files`,
	Run: func(cmd *cobra.Command, args []string) {
		generateViews()
	},
}

func generateViews() {
	createDirIfNotExist(output)
	if views == "..." {
		fmt.Printf("Genering all views at %v ...\n", output)
		makeViewsFiles()
	} else {
		vs := strings.Split(views, ",")
		fmt.Printf("Genering views:\n   %v\n\n", strings.Join(vs, ", \n   "))

		makeViewsFiles(vs...)
	}
}

func makeViewsFiles(views ...string) {
	names := make([]string, 0, len(TEMPLATES))
	if len(views) == 0 {
		for k := range TEMPLATES {
			names = append(names, k)
		}
	} else {
		names = views
	}
	for _, name := range names {
		var path string = path.Join(output, name+".html")
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if err := TEMPLATES[name].Execute(f, data.Error401); err != nil {
			panic(err)
		}
		fmt.Printf("generated %s.html\n", name)
	}
}

func createDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
