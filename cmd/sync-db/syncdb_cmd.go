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
	"fmt"
	"log"

	"github.com/medtune/go-utils/crypto"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/pkg/config"
	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/store/db"
	"github.com/medtune/beta-platform/pkg/store/model"
)

var (
	configFile  string
	createUsers bool
)

func init() {
	autoMigrateCmd.Flags().StringVarP(&configFile, "file", "f", "config.yml", "Configuration file name")
	autoMigrateCmd.Flags().BoolVarP(&createUsers, "create-users", "y", true, "Create default users before start")

	root.Cmd.AddCommand(autoMigrateCmd)
}

// syncdbCmd represents the syncdb command
var autoMigrateCmd = &cobra.Command{
	Use:     "automigrate",
	Aliases: []string{"sync-db", "syncdb"},
	Short:   "Auto migrate database",
	Long: `Sync database models by updating/creating existing 
database tables`,
	Run: func(cmd *cobra.Command, args []string) {
		autoMigrateDatabase()
	},
}

func autoMigrateDatabase() {
	var dbconfig *config.Database
	var usersConfig []*model.User

	if configuration, err := config.LoadConfigFromPath(configFile); err != nil {
		log.Fatal(err)
	} else {
		dbconfig = configuration.Database
		usersConfig = configuration.Create.Users
	}

	log.Printf("starting auto-migration on database: %v", dbconfig.Prod)

	cnxStr := db.ConnStr{
		Host:         dbconfig.Creds.Host,
		Database:     dbconfig.Prod,
		User:         dbconfig.Creds.User,
		Password:     dbconfig.Creds.Password,
		Port:         dbconfig.Creds.Port,
		SslMode:      dbconfig.SSLMode,
		MaxIdleConns: dbconfig.MIC,
		MaxOpenConns: dbconfig.MOC,
	}

	{
		engine, err := store.New(cnxStr)
		if err != nil {
			log.Fatal(err)
		}
		if err := engine.Sync(); err != nil {
			log.Fatal(err)
		}

		if createUsers && usersConfig != nil {
			if err := createUsersEngine(engine, usersConfig...); err != nil {
				log.Fatal(err)
			}
		}

	}

	log.Printf("\nstarting auto-migration on database: %v", dbconfig.Test)
	{
		cnxStr.Database = dbconfig.Test
		engine, err := store.New(cnxStr)
		if err != nil {
			log.Fatal(err)
		}
		if err := engine.Sync(); err != nil {
			log.Fatal(err)
		}

		if createUsers && usersConfig != nil {
			if err := createUsersEngine(engine, usersConfig...); err != nil {
				log.Fatal(err)
			}
		}

	}

}

func createUsersEngine(e *store.Store, us ...*model.User) error {
	for _, user := range us {
		if ok, err := e.Valid(user); err != nil || !ok {
			b, _ := yaml.Marshal(user)
			fmt.Printf("unvalid user data:\n\terror: %v\n\tuser:%s", err, string(b))
			continue
		}
		unhashed := user.Password
		user.Password = crypto.Sha256(user.Password)
		if _, err := e.Insert(user); err != nil {
			fmt.Printf("failed to create user: %s\n\terror: %v\n", user.Username, err)
			continue
		}
		fmt.Printf("created user %s %s\n", user.Email, user.Username)
		user.Password = unhashed
	}
	return nil
}
