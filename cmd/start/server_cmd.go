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

package start

import (
	"fmt"
	"log"
	"time"

	"github.com/medtune/beta-platform/hack/xlsx2pg/xlsx2pg"

	"github.com/gin-gonic/gin"
	"github.com/medtune/go-utils/crypto"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/medtune/beta-platform/cmd/root"
	"github.com/medtune/beta-platform/pkg/config"
	"github.com/medtune/beta-platform/pkg/initpkg"
	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/store/model"
	"github.com/medtune/beta-platform/server"
)

var (
	configFile    string
	port          int
	ginMode       int
	static        string
	staticBaseURL string
	syncdb        bool
	wait          bool
	maxattempts   int
	timestamp     int
	createUsers   bool
	cxpbaSync     bool
	cxpbaFile     string
	soft          bool
)

func init() {
	startCmd.Flags().IntVarP(&port, "port", "p", 8005, "port")
	startCmd.Flags().IntVarP(&ginMode, "gin-mode", "g", 0, "Gin server mode [0 OR 1]")
	startCmd.Flags().StringVarP(&configFile, "file", "f", "./config.yml", "Configuration file name")
	startCmd.Flags().StringVarP(&static, "static", "s", "./static", "Static files directory")
	startCmd.Flags().StringVarP(&staticBaseURL, "static-url", "z", "/static", "static base url")

	startCmd.Flags().BoolVarP(&soft, "soft", "S", false, "don't panic mode")
	startCmd.Flags().BoolVarP(&syncdb, "syncdb", "x", false, "Sync database before start")
	startCmd.Flags().BoolVarP(&createUsers, "create-users", "y", false, "Create default users before start")
	startCmd.Flags().BoolVarP(&wait, "wait", "w", false, "Wait all services to go up")
	startCmd.Flags().IntVarP(&maxattempts, "wait-attempts", "c", 60, "Wait max attempts")
	startCmd.Flags().IntVarP(&timestamp, "wait-timestamp", "t", 1, "Wait timestamp")

	startCmd.Flags().BoolVarP(&cxpbaSync, "sync-cxpba", "X", false, "Sync CXBPA before start")
	startCmd.Flags().StringVarP(&cxpbaFile, "cxpba-file", "F", "./CXPBA.xlsx", "CXPBA excel file name")

	root.Cmd.AddCommand(startCmd)
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"run", "run-server", "serve"},
	Short:   "Run Medtune beta server",
	Long:    `Run Medtune beta server`,
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func runServer() {
	// Set gin mode (0 for debug, 1 for production)
	if ginMode == 0 {
		gin.SetMode(gin.DebugMode)
	} else if ginMode == 1 {
		gin.SetMode(gin.ReleaseMode)
	} else if !soft {
		log.Fatalf("unknown gin mode: %v", ginMode)
	} else {

	}

	// Load configuration
	configuration, err := config.LoadConfigFromPath(configFile)
	if err != nil && !soft {
		log.Fatalf("failed to load configuration: %v\n\t%v\n", err, configFile)
	}

	// Init packages
	if err := initpkg.InitFromConfig(configuration); err != nil && !soft {
		log.Fatalf("failed to initialize packages: %v\n\t%v\n\t%v\n", err, configFile, configuration)
	}

	// Syndb
	// Create tables from models structs
	if syncdb {
		fmt.Printf("trying to connect database %s\n", configuration.Database.Prod)

		var err error

		// First sync
		err = store.Agent.Sync()
		if err != nil {

			if err != nil && wait {
				attempt := 0
				// waiting for err == 0 or attempt > maxattempts
				for err != nil && attempt < maxattempts {
					time.Sleep(time.Duration(timestamp) * time.Second)
					fmt.Printf("waiting for database response (host: %s) (attempt: %d)\n", configuration.Database.Creds.Host, attempt)
					err = store.Agent.Sync()
					attempt++
				}
			}
			if err != nil && !soft {
				log.Fatalln(err)
			}

			fmt.Println("connected to database...")
		}

		// Create config.Users
		if createUsers && configuration.Create != nil {
			createUsersEngine(store.Agent, configuration.Create.Users...)
		}
	}

	if cxpbaSync {
		if err := xlsx2pg.SyncCXPBAexcel(cxpbaFile, configFile); err != nil {
			log.Println(err)
		}
	}

	if configuration.Public.Static != "" {
		static = configuration.Public.Static
	}

	if configuration.Public.Prefix != "" {
		staticBaseURL = configuration.Public.Prefix
	}

	fmt.Println("starting server...")

	Server := server.New(
		static,
		staticBaseURL,
		port,
	)

	if err := Server.Run(); err != nil {
		log.Fatalln(err)
	}
}

func createUsersEngine(e *store.Store, us ...*model.User) error {
	for _, user := range us {
		if ok, err := e.Valid(user); err != nil || !ok {
			b, _ := yaml.Marshal(user)
			fmt.Printf("unvalid user data:\n\terror: %v\n\tuser:%s", err, string(b))
			continue
		}

		unhashed := (*user).Password
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
