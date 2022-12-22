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
	gorav := &goravel.Goravel{}
	err = gorav.New(path)
	if err != nil {
		log.Fatal(err)
	}
	gorav.AppName = "testapp"

	gorav.InfoLog.Println("Debug is set to", gorav.Debug)

	app := &application{
		App: gorav,
	}
	return app
}
