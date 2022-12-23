package main

import (
	"goravel"
	"log"
	"os"
	handlers "testapp/handlers"
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

	routeHandlers := &handlers.Handlers{
		App: gorav,
	}

	app := &application{
		App:      gorav,
		Handlers: routeHandlers,
	}

	app.App.Routes = app.routes()
	return app
}
