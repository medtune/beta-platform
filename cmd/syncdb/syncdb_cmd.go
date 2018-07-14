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

package syncdb

import (
	"log"

	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/pkg/config"
	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/store/db"
	"github.com/spf13/cobra"
)

var (
	configFile string
)

func init() {
	autoMigrateCmd.Flags().StringVarP(&configFile, "file", "f", "config.yaml", "Configuration file name")

	root.Cmd.AddCommand(autoMigrateCmd)
}

// syncdbCmd represents the syncdb command
var autoMigrateCmd = &cobra.Command{
	Use:     "automigrate",
	Aliases: []string{"syncdb"},
	Short:   "Auto migrate database",
	Long: `Sync database models by updating/creating existing 
database tables`,
	Run: func(cmd *cobra.Command, args []string) {
		autoMigrateDatabase()
	},
}

func autoMigrateDatabase() {
	var dbconfig *config.Database

	if cfg, err := config.LoadConfigFromPath(configFile); err != nil {
		log.Panic(err)
	} else {
		dbconfig = cfg.Database
	}

	log.Printf("starting auto-migration on database: %v", dbconfig.Prod)

	{
		engine, err := store.New(db.ConnStr{
			Host:     dbconfig.Creds.Host,
			Database: dbconfig.Prod,
			User:     dbconfig.Creds.User,
			Password: dbconfig.Creds.Password,
			Port:     dbconfig.Creds.Port,
			SslMode:  0,
			Maxconn:  2,
		})
		if err != nil {
			log.Panic(err)
		}
		if err := engine.Sync(); err != nil {
			log.Panic(err)
		}

	}

	log.Printf("starting auto-migration on database: %v", dbconfig.Test)
	{
		engine, err := store.New(db.ConnStr{
			Host:     dbconfig.Creds.Host,
			Database: dbconfig.Test,
			User:     dbconfig.Creds.User,
			Password: dbconfig.Creds.Password,
			Port:     dbconfig.Creds.Port,
			SslMode:  0,
			Maxconn:  2,
		})
		if err != nil {
			log.Panic(err)
		}
		if err := engine.Sync(); err != nil {
			log.Panic(err)
		}

	}

}
