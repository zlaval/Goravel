package goravel

import "net/http"

func (g *Goravel) SessionLoad(next http.Handler) http.Handler {
	g.InfoLog.Println("SessionLoad called")
	return g.Session.LoadAndSave(next)
}
