package main

import (
	"fmt"

	"github.com/medtune/beta-platform/pkg/config"
	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/store/db"
	"github.com/medtune/beta-platform/pkg/store/model"
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

	// Get data
	analysisData := make([]*model.PathologyAnalysisLevel, 0, 100)
	specs := make([]*model.SpecAnalysisPool, 0, 100)

	var sheetN *xlsx.Sheet
	var sheetMA *xlsx.Sheet

	for _, sheet := range xlFile.Sheets {
		if sheet.Name == "Norms" {
			sheetN = sheet
		} else if sheet.Name == "Medical Analysis" {
			sheetMA = sheet
		}
	}

	// Copy medical analysis
	for index, raw := range sheetMA.Rows[1:] {
		analysisData = append(analysisData,
			&model.PathologyAnalysisLevel{
				int64(index),
				raw.Cells[0].String(),
				raw.Cells[1].String(),
				raw.Cells[2].String(),
				raw.Cells[3].String(),
				raw.Cells[4].String(),
				raw.Cells[5].String(),
				raw.Cells[6].String(),
				raw.Cells[7].String(),
				raw.Cells[8].String(),
				raw.Cells[9].String(),
				raw.Cells[10].String(),
				raw.Cells[11].String(),
				raw.Cells[12].String(),
				raw.Cells[13].String(),
				raw.Cells[14].String(),
				raw.Cells[16].String(),
				raw.Cells[17].String(),
				raw.Cells[18].String(),
				raw.Cells[19].String(),
				raw.Cells[20].String(),
				raw.Cells[21].String(),
				raw.Cells[22].String(),
				raw.Cells[23].String(),
				raw.Cells[24].String(),
				raw.Cells[25].String(),
			},
		)
	}

	for _, e := range analysisData {
		fmt.Println(e.Name)
	}

	// Copy medical analysis
	for index, raw := range sheetN.Rows[1:] {
		min, _ := raw.Cells[2].Int64()
		max, _ := raw.Cells[3].Int64()
		specs = append(specs,
			&model.SpecAnalysisPool{
				int64(index),
				raw.Cells[0].String(),
				raw.Cells[1].String(),
				min,
				max,
			},
		)
	}
	for _, e := range specs {
		fmt.Println(e.Name)
	}

}
