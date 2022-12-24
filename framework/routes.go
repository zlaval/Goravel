package goravel

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func (g *Goravel) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Recoverer)
	mux.Use(g.SessionLoad)
	if g.Debug {
		mux.Use(middleware.Logger)
	}
	//mux.Get("/", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprint(writer, "Welcome to Goravel")
	//})
	return mux
}
