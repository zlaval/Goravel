package main

import (
	"goravel"
	"testapp/handlers"
)

type application struct {
	App      *goravel.Goravel
	Handlers *handlers.Handlers
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
