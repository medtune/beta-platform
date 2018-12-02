package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/medtune/beta-platform/internal/config"
	"github.com/spf13/cobra"
)

var (
	versionToFind string
	versionNew    string

	configs          string
	makefiles        string
	readmes          string
	nextMajor        bool
	nextMinor        bool
	nextPatch        bool
	initSubVersions  bool
	forceVersion     string
	interruptOnError bool
)

func init() {
	Cmd.Flags().StringVarP(&configs, "configs", "C", "config.yml", "configuration files 'config.yml'")
	Cmd.Flags().StringVarP(&makefiles, "makefiles", "K", "Makefile", "makefile files")
	Cmd.Flags().StringVarP(&readmes, "readmes", "R", "README.md", "readme files")
	Cmd.Flags().BoolVarP(&nextPatch, "patch", "p", false, "")
	Cmd.Flags().BoolVarP(&nextMinor, "minor", "m", false, "")
	Cmd.Flags().BoolVarP(&nextMajor, "major", "M", false, "")
	Cmd.Flags().BoolVarP(&initSubVersions, "reinit-subs", "r", true, "")
	Cmd.Flags().StringVarP(&forceVersion, "", "f", "", "")
	Cmd.Flags().BoolVarP(&interruptOnError, "interupt", "i", false, "")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "bump-version",
	Short: "tool to bump versions accross project confs & packages",
	Run: func(cmd *cobra.Command, args []string) {
		if configs != "" {
			files := strings.Split(configs, ",")
			for _, file := range files {
				updateVersion(file, "config")
			}
		}

		if makefiles != "" {
			files := strings.Split(makefiles, ",")
			for _, file := range files {
				updateVersion(file, "makefile")
			}
		}

		if readmes != "" {
			files := strings.Split(readmes, ",")
			for _, file := range files {
				updateVersion(file, "readme")
			}
		}
	},
}

func main() {
	if err := Cmd.Execute(); err != nil {
		fmt.Printf("[FATAL] %v", err)
		os.Exit(1)
	}
}

func updateVersion(target string, t string) {
	switch t {
	case "config":
		updateStartupConfigVersion(target)
	case "makefile":
		updateMakefileVersion(target)
	case "readme":
		updateReadmeVersion(target)
	}
}

func updateStartupConfigVersion(target string) error {
	fmt.Printf("config %s\n", target)
	cfg, err := config.LoadConfigFromPath(target)
	if err != nil {
		return err
	}

	cfg.Meta.Version = versionNew

	return nil
}

func updateMakefileVersion(target string) error {
	fmt.Printf("makefile %s\n", target)

	fb, err := ioutil.ReadFile(target)
	if err != nil {
		return err
	}

	b := bytes.NewBuffer(fb)
	buf := bufio.NewReader(b)

	fmt.Println(buf.ReadLine())
	fmt.Println(buf.ReadLine())
	fmt.Println(buf.ReadLine())
	fmt.Println(buf.ReadLine())

	return nil
}

func updateReadmeVersion(target string) {
	fmt.Printf("readme %s\n", target)
}
