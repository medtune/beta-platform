package main

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/config"
	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/store/db"
	"github.com/tealeg/xlsx"
)

// Sync database with CXPBA (Chest Xray pathology biological analysis) data
func syncCXPBAexcel(file string) {
	var dbconfig *config.Database

	if configuration, err := config.LoadConfigFromPath(configFile); err != nil {
		panic(err)
	} else {
		dbconfig = configuration.Database
	}

	// create engine
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
		panic(err)
	}

	// Engine sync
	err = engine.Sync()
	if err != nil {
		panic(err)
	}

	// Read XLSXFILE
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		panic(err)
	}

	// Dump data
	for _, sheet := range xlFile.Sheets {
		fmt.Println("-------", sheet)
		for _, row := range sheet.Rows {
			fmt.Println("------", row)
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
