package main

import "goravel"

type application struct {
	App *goravel.Goravel
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
