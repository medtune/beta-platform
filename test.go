package main

import (
	"fmt"
	"log"

	"github.com/medtune/beta-platform/pkg/store"
	"github.com/medtune/beta-platform/pkg/store/db"
	"github.com/medtune/beta-platform/pkg/store/model"
)

func main() {
	engine, err := store.New(db.ConnStr{
		Host:     "localhost",
		Database: "medtune",
		User:     "amine",
		Password: "a",
		Port:     5432,
		SslMode:  0,
		Maxconn:  2,
	})
	if err != nil {
		log.Panic(err)
	}

	engine.Sync2(&model.User{})
	err = engine.CreateUser("amine", "pw")
	if err != nil {
		log.Panic(err)
	}
	user, err := engine.GetUser("amine")
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(user)
}
