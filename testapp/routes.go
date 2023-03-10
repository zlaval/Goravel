package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (a *application) routes() *chi.Mux {

	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions", a.Handlers.SessionTest)
	a.App.Routes.Get("/dbtest", func(w http.ResponseWriter, r *http.Request) {
		query := "select id, name from users where id = 1"
		row := a.App.DB.Pool.QueryRowContext(r.Context(), query)

		var id int
		var name string
		err := row.Scan(&id, &name)
		if err != nil {
			a.App.ErrorLog.Println(err)
		}
		fmt.Fprintf(w, "%d %s", id, name)
	})

	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))
	return a.App.Routes
}
