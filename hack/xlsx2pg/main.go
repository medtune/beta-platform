package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/medtune/beta-platform/hack/xlsx2pg/xlsx2pg"
)

var (
	configf string
	target  string
)

func init() {
	cmd.Flags().StringVarP(&target, "target", "t", "./CXPBA.xlsx", "CXPBA file name")
	cmd.Flags().StringVarP(&configf, "config", "f", "./dev.config.yml", "Configuration file name")
}

var cmd = &cobra.Command{
	Use:   "xlsx2pg",
	Short: "Sync CXPBA excel file with pg database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic(fmt.Errorf("Need one argument (data file name)"))
		}
		if err := xlsx2pg.SyncCXPBAexcel(target, configf); err != nil {
			log.Fatal(err)
		}
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
