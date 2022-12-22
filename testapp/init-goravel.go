package main

import (
	"goravel"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	g := &goravel.Goravel{}
	err = g.New(path)
	if err != nil {
		log.Fatal(err)
	}
	g.AppName = "testapp"
	g.Debug = true
	app := &application{
		App: g,
	}
	return app
}
