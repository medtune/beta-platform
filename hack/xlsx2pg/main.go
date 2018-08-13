package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

func init() {
	cmd.Flags().StringVarP(&configFile, "config", "f", "./config.yml", "Configuration file name")
}

var cmd = &cobra.Command{
	Use:   "xlsx2pg",
	Short: "Sync CXPBA excel file with pg database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic(fmt.Errorf("Need one argument (data file name)"))
		}
		syncCXPBAexcel(args[0])
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
